package resolver

import (
	"context"
	"fmt"
	"room-reservation/ent"
	generated "room-reservation/graphql"
	"room-reservation/internal/util"
	"room-reservation/service"

	"github.com/99designs/gqlgen/graphql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client               *ent.Client
	validator            *validator.Validate
	validationTranslator ut.Translator
	logger               *zap.Logger
	serviceRegistry      service.Service
}

func NewSchema(serviceRegistry service.Service, client *ent.Client, validator *validator.Validate, validationTranslator ut.Translator, logger *zap.Logger) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:               client,
			validator:            validator,
			validationTranslator: validationTranslator,
			logger:               logger,
			serviceRegistry:      serviceRegistry,
		},
	})
}

func (r *Resolver) ValidationResolver() func(ctx context.Context, obj interface{}, next graphql.Resolver, contrains string) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, contrains string) (interface{}, error) {
		val, err := next(ctx)
		if err != nil {
			r.logger.Error("Getting error when extract values from context", zap.Error(err))
			return nil, util.WrapGQLInternalError(ctx)
		}

		fieldName := *graphql.GetPathContext(ctx).Field

		err = r.validator.Var(val, contrains)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return nil, fmt.Errorf("%s:%s", fieldName, validationErrors[0].Translate(r.validationTranslator))
		}

		return val, nil
	}
}
