package database

import "go.mongodb.org/mongo-driver/mongo"

var client *mongo.Client
var zonesCollection *mongo.Collection
var COLLECTION = "location"
