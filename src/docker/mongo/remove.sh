#!/bin/bash

# Removing registers with min value on database sqlite by robots....
/usr/bin/mongo gerbo --eval 'db.movies.remove({"_id": db.movies.find({ "id": db.movies.aggregate([{ $group: { _id: {}, min: { $min: "$id" } } }]).toArray()[0].min }).toArray()[0]._id });'