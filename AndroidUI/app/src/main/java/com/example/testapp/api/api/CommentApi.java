package com.example.testapp.api.api;

import com.example.testapp.api.models.Comment;
import com.example.testapp.api.models.Station;

import java.util.List;

import io.reactivex.Single;
import okhttp3.ResponseBody;
import retrofit2.Response;
import retrofit2.http.Body;
import retrofit2.http.DELETE;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.PUT;
import retrofit2.http.Path;
import retrofit2.http.Query;

public interface CommentApi {
    @POST("stations/{sid}/comments")
    public Single<Response<Comment>> createComment(@Path("sid") String sid, @Body Comment comment);

    @GET("stations/{sid}/comments/read")
    public Single<Response<List<Comment>>> readComments(@Path("sid") String sid, @Query("skip") int skip, @Query("limit") int limit);

    @GET("stations/{sid}/comments/{id}")
    public Single<Response<Comment>> findById(@Path("sid") String sid, @Path("id") String id);

    @PUT("stations/{sid}/comments/{id}")
    public Single<Response<Comment>> updateById(@Path("sid") String sid, @Path("id") String id, @Body Comment comment);

    @DELETE("stations/{sid}/comments/{id}")
    public Single<ResponseBody> deleteById(@Path("sid") String sid, @Path("id") String id);
}
