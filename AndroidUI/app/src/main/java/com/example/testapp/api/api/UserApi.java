package com.example.testapp.api.api;

import com.example.testapp.api.models.User;

import io.reactivex.Single;
import okhttp3.ResponseBody;
import retrofit2.Response;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Path;

public interface UserApi {
    @POST("users")
    public Single<Response<User>> createUser(@Body User body);

    @POST("login")
    public Single<Response<User>> login(@Body User body);

    @GET("logout/{id}")
    public Single<ResponseBody> logout(@Path("id") String id);
}
