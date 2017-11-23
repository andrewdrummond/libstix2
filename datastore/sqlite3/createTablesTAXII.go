// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllTAXIITables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllTAXIITables() {
	ds.createTAXIITable(datastore.DB_TABLE_TAXII_COLLECTION_CONTENT, ds.collectionContent())
	ds.createTAXIITable(datastore.DB_TABLE_TAXII_COLLECTION, ds.collection())
	ds.createTAXIITable(datastore.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE, ds.collectionMediaType())
}

// ----------------------------------------------------------------------
//
// Private Methods
//
// These methods return a list of fields that is used for creating the
// database table.
//
// ----------------------------------------------------------------------

/*
collectionContent - This method will return the properties that make up the
collection content table

date_added    = The date that this object was added to the collection
collection_id = The collection ID that this object is tied to
stix_id       = The STIX ID for the object that is being mapped to a collection.
  We do not use the object_id here or the row_id as that would point to a
  specific version and we need to be able to find all versions of an object.
  and if we used row_id for example, it would require two queries, the first
  to get the SITX ID and then the second to get all objects with that STIX ID.
*/
func (ds *Sqlite3DatastoreType) collectionContent() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"date_added" TEXT NOT NULL,
 	"collection_id" TEXT NOT NULL,
 	"stix_id" TEXT NOT NULL
 	`
}

/*
collection - This method will return the properties that make up the collection
table

date_added  = The date that this collection was added to the system
enabled     = Is this collection currently enabled
id 		    = The collection ID, a UUIDv4 value
title 	    = The title of this collection
description = A long description about this collection
can_read    = A boolean flag that indicates if one can read from this collection
can_write   = A boolean flag that indicates if one can write to this collection
*/
func (ds *Sqlite3DatastoreType) collection() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"date_added" TEXT NOT NULL,
	"enabled" INTEGER(1,0) NOT NULL DEFAULT 1,
	"id" TEXT NOT NULL,
	"title" TEXT NOT NULL,
	"description" TEXT,
	"can_read" INTEGER(1,0) NOT NULL DEFAULT 0,
	"can_write" INTEGER(1,0) NOT NULL DEFAULT 0
	`
}

/*
collectionMediaType  - This method will return the properties that make up the
collection media type table

collection_id = The collection ID, a UUIDv4 value
media_type    = The media types supported on this collection
*/
func (ds *Sqlite3DatastoreType) collectionMediaType() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"collection_id" TEXT NOT NULL,
	"media_type" TEXT NOT NULL
	`
}