package com.example.testapp.api.models;

import com.google.gson.annotations.SerializedName;

import java.time.LocalDateTime;
import java.util.Date;

public class Comment {
    @SerializedName("_id")
    private String id;
    @SerializedName("user_id")
    private String userID;
    @SerializedName("user_name")
    private String userName;
    @SerializedName("text")
    private String text;
    @SerializedName("rating")
    private float rating;
    @SerializedName("create_at")
    private String createAt;
    @SerializedName("update_at")
    private String updateAt;

    public Comment() {
    }

    public Comment(String id, String userID, String userName, String text, float raiting, String createAt, String updateAt) {
        this.id = id;
        this.userID = userID;
        this.userName = userName;
        this.text = text;
        this.rating = raiting;
        this.createAt = createAt;
        this.updateAt = updateAt;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUserID() {
        return userID;
    }

    public void setUserID(String userID) {
        this.userID = userID;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public float getRating() {
        return rating;
    }

    public void setRating(float rating) {
        this.rating = rating;
    }

    public String getCreateAt() {
        return createAt;
    }

    public void setCreateAt(String createAt) {
        this.createAt = createAt;
    }

    public String getUpdateAt() {
        return updateAt;
    }

    public void setUpdateAt(String updateAt) {
        this.updateAt = updateAt;
    }
}
