package com.example.testapp.api.api;

import com.example.testapp.api.models.RouteResponse;

import io.reactivex.Observable;
import retrofit2.Response;
import retrofit2.http.GET;
import retrofit2.http.Query;

public interface MapsApi {

    @GET("/maps/api/directions/json")
    Observable<Response<RouteResponse>> getRoute(
            @Query(value = "origin", encoded = false) String position,
            @Query(value = "destination", encoded = false) String destination,
            @Query("sensor") boolean sensor,
            @Query("language") String language);
}
