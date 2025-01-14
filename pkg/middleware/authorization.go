package middleware

import (
	"fmt"
	"net/http"

	"github.com/qingwave/weave/pkg/authorization"
	"github.com/qingwave/weave/pkg/common"
	"github.com/qingwave/weave/pkg/model"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := common.GetUser(c)
		if user == nil {
			user = &model.User{}
		}

		ri := common.GetRequestInfo(c)
		if ri == nil {
			common.ResponseFailed(c, http.StatusBadRequest, fmt.Errorf("failed to get request info"))
			c.Abort()
			return
		}

		if ri.IsResourceRequest {
			resource := ri.Resource
			ok := authorization.Enforce(user.Name, ri.Namespace, resource, ri.Name, ri.Verb)
			if !ok {
				if user.Name == "" {
					common.ResponseFailed(c, http.StatusUnauthorized, nil)
				} else {
					common.ResponseFailed(c, http.StatusForbidden, fmt.Errorf("user [%s] is forbidden for resource %s in namespace %s", user.Name, resource, ri.Namespace))
				}
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
