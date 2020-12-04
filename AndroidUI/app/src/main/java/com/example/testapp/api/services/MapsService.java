package com.example.testapp.api.services;

import com.example.testapp.App;
import com.example.testapp.Helper;
import com.example.testapp.api.api.MapsApi;
import com.example.testapp.api.api.UserApi;
import com.example.testapp.api.models.User;

import java.io.IOException;

import okhttp3.Interceptor;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.logging.HttpLoggingInterceptor;
import retrofit2.Retrofit;
import retrofit2.adapter.rxjava2.RxJava2CallAdapterFactory;
import retrofit2.converter.gson.GsonConverterFactory;

public class MapsService {

    private MapsApi mapsApi;

    public MapsService(){
        Retrofit retrofit = createRetrofit();
        mapsApi = retrofit.create(MapsApi.class);
    }

    private OkHttpClient createOkHttpClient() {
        OkHttpClient.Builder httpClient = new OkHttpClient.Builder();
        httpClient.addInterceptor(new Interceptor() {
            @Override
            public Response intercept(Chain chain) throws IOException {
                Request request = chain.request().newBuilder()
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
                .baseUrl("https://maps.googleapis.com")
                .client(createOkHttpClient())
                .addCallAdapterFactory(RxJava2CallAdapterFactory.create())
                .addConverterFactory(GsonConverterFactory.create())
                .build();
    }

    public MapsApi getMapsApi() {
        return mapsApi;
    }

    public void setMapsApi(MapsApi mapsApi) {
        this.mapsApi = mapsApi;
    }
}
