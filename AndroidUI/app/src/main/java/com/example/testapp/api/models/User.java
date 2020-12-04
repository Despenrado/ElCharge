package com.example.testapp.api.models;

import com.google.gson.annotations.SerializedName;

import java.text.SimpleDateFormat;
import java.time.LocalDateTime;
import java.util.Date;


public class User {
    @SerializedName("_id")
    private String id;
    @SerializedName("user_name")
    private String userName;
    @SerializedName("email")
    private String email;
    @SerializedName("password")
    private String password;
    @SerializedName("update_at")
    private String updateAt;

    public User() {
    }

    public User(String id, String userName, String email, String password, String updateAt) {
        this.id = id;
        this.userName = userName;
        this.email = email;
        this.password = password;
        this.updateAt = updateAt;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getUpdateAt() {
        return updateAt;
    }

    public void setUpdateAt(String updateAt) {
        this.updateAt = updateAt;
    }
}
