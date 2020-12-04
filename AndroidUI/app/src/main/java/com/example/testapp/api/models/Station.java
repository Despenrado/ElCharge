package com.example.testapp.api.models;

import com.google.gson.annotations.SerializedName;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Date;

public class Station {
    @SerializedName("_id")
    private String id;
    @SerializedName("owner_id")
    private String ownerID;
    @SerializedName("description")
    private String description;
    @SerializedName("station_name")
    private String stationName;
    @SerializedName("latitude")
    private double latitude;
    @SerializedName("longitude")
    private double longitude;
    @SerializedName("update_at")
    private String updateAt;
    @SerializedName("rating")
    private float rating;
    @SerializedName("comments")
    private ArrayList<Comment> comments;

    public Station() {
    }

    public Station(String id, String ownerID, String description, String stationName, double latitude, double longitude, String updateAt, float rating, ArrayList<Comment> comments) {
        this.id = id;
        this.ownerID = ownerID;
        this.description = description;
        this.stationName = stationName;
        this.latitude = latitude;
        this.longitude = longitude;
        this.updateAt = updateAt;
        this.rating = rating;
        this.comments = comments;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public float getRating() {
        return rating;
    }

    public void setRating(float rating) {
        this.rating = rating;
    }

    public ArrayList<Comment> getComments() {
        return comments;
    }

    public void setComments(ArrayList<Comment> comments) {
        this.comments = comments;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getStationName() {
        return stationName;
    }

    public void setStationName(String stationName) {
        this.stationName = stationName;
    }

    public String getUpdateAt() {
        return updateAt;
    }

    public double getLatitude() {
        return latitude;
    }

    public void setLatitude(double latitude) {
        this.latitude = latitude;
    }

    public double getLongitude() {
        return longitude;
    }

    public void setLongitude(double longitude) {
        this.longitude = longitude;
    }

    public void setUpdateAt(String updateAt) {
        this.updateAt = updateAt;
    }

    public String getOwnerID() {
        return ownerID;
    }

    public void setOwnerID(String ownerID) {
        this.ownerID = ownerID;
    }
}
