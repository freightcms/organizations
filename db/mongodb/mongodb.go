package mongodb

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"slices"

	"github.com/freightcms/organizations/db"
	"github.com/freightcms/organizations/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrganizationResourceManagerContextKey string

const (
	// ContextKey used to fetch or put the Person Resource Manager into the context
	ContextKey OrganizationResourceManagerContextKey = "organizationResourceManagerContextKey"
)

type resourceManager struct {
	session mongo.SessionContext
}

// WithContext fetches the mongo db session context from that passed argument (parent context)
// ,appends the Organization manager and returns all with the new context.
func WithContext(session mongo.SessionContext) context.Context {
	if session == nil {
		panic("Could not fetch session from context")
	}
	mgr := NewOrganizationManager(session)
	return context.WithValue(session, ContextKey, mgr)
}

func FromContext(ctx context.Context) db.OrganizationResourceManager {
	val := ctx.Value(ContextKey)
	if val == nil {
		panic(errors.New("could not fetch OrganizationResourceManager from context"))
	}

	return val.(*resourceManager)
}

// Get implements db.PersonResourceManager.
func (r *resourceManager) Get(query *db.OrganizationQuery) ([]*models.Organization, error) {
	projection := bson.D{}

	// see https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/project/
	for _, fieldName := range query.Fields {
		// for security reasons we only want people to be able to query the objects that they should be able to
		if slices.Contains([]string{"id", "dba", "name", "rollupId", "mailingAddress", "billingAddress"}, fieldName) {
			projection = append(projection, bson.E{
				Key:   fieldName,
				Value: 1,
			})
		}
	}
	if len(query.SortBy) != 0 {
		if !slices.Contains([]string{"_id", "id"}, query.SortBy) {
			return nil, fmt.Errorf("%s is not a valid sortBy option", query.SortBy)
		}
	}
	sort := bson.D{bson.E{Key: query.SortBy, Value: 1}}
	opts := options.Find().
		SetSort(sort).
		SetLimit(int64(query.PageSize)).
		SetSkip(int64((query.Page) * query.PageSize)).
		SetProjection(projection)

	cursor, err := r.collection().Find(r.session, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	results := []*models.Organization{}
	for cursor.Next(r.session) {
		var result models.Organization
		if err := cursor.Decode(&result); err != nil {
			fmt.Printf("Error occured fetching Organization record %s\n", err.Error())
			continue
		}
		results = append(results, &result)
	}
	return results, nil

}

func (r *resourceManager) CreateOrganization(model *models.Organization) (interface{}, error) {
	if model.MailingAddress != nil {
		model.MailingAddress.Id = primitive.NewObjectID()
	}
	insertedResult, err := r.collection().InsertOne(r.session,
		&bson.M{
			"dba":            model.DBA,
			"name":           model.Name,
			"rollupId":       model.RollupID,
			"mailingAddress": model.MailingAddress,
			"billingAddress": model.BillingAddress,
		},
		options.InsertOne(),
	)
	if err != nil {
		return nil, err
	}
	if model.RollupID != nil {
		rollupId, err := primitive.ObjectIDFromHex(*model.RollupID)
		if err != nil {
			return nil, err
		}

		count, err := r.collection().CountDocuments(r.session, bson.M{"_id": rollupId})
		if err != nil {
			return nil, err
		}
		if count == 0 {
			return nil, errors.New("cannot set rollup id to id that does not match another organization")
		}
	}
	id := insertedResult.InsertedID.(primitive.ObjectID)

	return id.Hex(), nil
}

func (r *resourceManager) DeleteOrganization(id interface{}) error {
	if reflect.TypeOf(id).Kind() != reflect.String {
		return fmt.Errorf("cannot use typeof %s as id parameter", reflect.TypeOf(id).String())
	}

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}

	if err = r.session.StartTransaction(); err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	if _, err = r.collection().DeleteOne(r.session, filter); err != nil {
		return err
	}
	if _, err = r.collection().UpdateMany(r.session, bson.M{"rollupId": objectId}, &bson.M{"rollupId": ""}); err != nil {
		return err
	}
	err = r.session.CommitTransaction(r.session)
	return err
}

func (r *resourceManager) GetById(id interface{}) (*models.Organization, error) {
	var result models.Organization

	if reflect.TypeOf(id).Kind() != reflect.String {
		return nil, fmt.Errorf("cannot use typeof %s as id parameter", reflect.TypeOf(id).String())
	}

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectId}
	if err := r.collection().FindOne(r.session, &filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *resourceManager) UpdateOrganization(id interface{}, model *models.Organization) error {
	if reflect.TypeOf(id).Kind() != reflect.String {
		return fmt.Errorf("cannot use typeof %s as id parameter", reflect.TypeOf(id).String())
	}

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}

	if model.RollupID != nil {
		rollupId, err := primitive.ObjectIDFromHex(*model.RollupID)
		if err != nil {
			return err
		}

		count, err := r.collection().CountDocuments(r.session, bson.M{"_id": rollupId})
		if err != nil {
			return err
		}
		if count == 0 {
			return errors.New("cannot set rollup id to id that does not match another organization")
		}
	}
	filter := bson.M{"_id": objectId}
	result, err := r.collection().UpdateOne(r.session, filter, model)

	if result.MatchedCount == 0 {
		return fmt.Errorf("could not find Organization with id %s", id)
	}
	return err
}

func (r *resourceManager) collection() *mongo.Collection {
	coll := r.session.Client().Database("freightcms").Collection("organizations")
	return coll
}

func NewOrganizationManager(session mongo.SessionContext) db.OrganizationResourceManager {
	return &resourceManager{session: session}
}
