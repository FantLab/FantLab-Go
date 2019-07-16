#!/bin/bash

PACKAGE=$1

if [ -z $PACKAGE ]; then
    echo "Usage: make_api_module.sh <package_name>"
    exit 1
fi

cd app
cd modules
mkdir -p $PACKAGE && cd $PACKAGE

if [ ! -f "datacontroller.go" ]; then
	echo "package $PACKAGE" > "datacontroller.go"
fi

if [ ! -f "api.go" ]; then
cat > api.go << EOF
package $PACKAGE

import (
	"fantlab/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) DeleteMe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"package": "$PACKAGE"})
}
EOF
fi

echo "API module $PACKAGE has been successfully created!"
