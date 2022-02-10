package com.example.testapp.api.services;

import com.example.testapp.App;
import com.example.testapp.utils.Helper;
import com.example.testapp.api.api.CommentApi;
import com.example.testapp.api.api.StationApi;
import com.example.testapp.api.api.UserApi;
import com.example.testapp.models.User;

import java.io.IOException;
import okhttp3.Interceptor;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.logging.HttpLoggingInterceptor;
import retrofit2.Retrofit;
import retrofit2.adapter.rxjava2.RxJava2CallAdapterFactory;
import retrofit2.converter.gson.GsonConverterFactory;

public class ElchargeService {
    String apiAddr;
    String token;
    UserApi userApi;
    StationApi stationApi;
    CommentApi commentApi;
    User user;

    public ElchargeService(){
        apiAddr = Helper.getConfigValue(App.getAppContext(),"apiserver_addr");
        token = Helper.readFromSharedPreferences(App.getAppContext(),"jwt_token");
        user = new User();
        user.setId(Helper.readFromSharedPreferences(App.getAppContext(),"user_id"));
        user.setUserName(Helper.readFromSharedPreferences(App.getAppContext(),"user_name"));
        Retrofit retrofit = createRetrofit();
        userApi = retrofit.create(UserApi.class);
        stationApi = retrofit.create(StationApi.class);
        commentApi = retrofit.create(CommentApi.class);
    }

    private OkHttpClient createOkHttpClient() {
        OkHttpClient.Builder httpClient = new OkHttpClient.Builder();
        httpClient.addInterceptor(new Interceptor() {
            @Override
            public Response intercept( Chain chain) throws IOException {
                Request request = chain.request().newBuilder()
                        .addHeader("Authorization", token)
                        .build();
                return chain.proceed(request);
            }
        });
        HttpLoggingInterceptor logging = new HttpLoggingInterceptor();
        logging.setLevel(HttpLoggingInterceptor.Level.BODY);
        httpClient.addInterceptor(logging);
        return httpClient.build();
    }

    private Retrofit createRetrofit() {
        return new Retrofit.Builder()
                .baseUrl(apiAddr)
                .client(createOkHttpClient())
                .addCallAdapterFactory(RxJava2CallAdapterFactory.create())
                .addConverterFactory(GsonConverterFactory.create())
                .build();
    }

    public String getApiAddr() {
        return apiAddr;
    }

    public void setApiAddr(String apiAddr) {
        this.apiAddr = apiAddr;
    }

    public UserApi getUserApi() {
        return userApi;
    }

    public void setUserApi(UserApi userApi) {
        this.userApi = userApi;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        Helper.saveToSharedPreferences(App.getAppContext(),"user_id", user.getId());
        Helper.saveToSharedPreferences(App.getAppContext(),"user_name", user.getUserName());
        this.user = user;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        Helper.saveToSharedPreferences(App.getAppContext(),"jwt_token", token);
        this.token = token;
    }

    public CommentApi getCommentApi() {
        return commentApi;
    }

    public void setCommentApi(CommentApi commentApi) {
        this.commentApi = commentApi;
    }

    public StationApi getStationApi() {
        return stationApi;
    }

    public void setStationApi(StationApi stationApi) {
        this.stationApi = stationApi;
    }
}
