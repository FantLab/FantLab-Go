#!/bin/bash

PACKAGE=$1

if [ -z $PACKAGE ]; then
    echo "Usage: make_module.sh package_name"
    exit 1
fi

cd modules
mkdir -p $PACKAGE && cd $PACKAGE
for FILE in datacontroller dbcontroller dbmodels models; do
    if [ ! -f "$FILE.go" ]; then
        echo "package $PACKAGE" > "$FILE.go"
    fi
done

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

echo "Module $PACKAGE has been successfully created!"
