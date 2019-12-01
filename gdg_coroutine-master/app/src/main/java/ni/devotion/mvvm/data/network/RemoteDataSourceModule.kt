package ni.devotion.mvvm.data.network

import ni.devotion.mvvm.data.network.`interface`.DepartmentsInterface
import ni.devotion.mvvm.data.network.adapters.CoroutineCallAdapterFactory
import ni.devotion.mvvm.data.network.headerInterceptor.HeaderInterceptor
import ni.devotion.mvvm.BuildConfig
import ni.devotion.mvvm.data.network.`interface`.RestaurantsInterface
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import org.koin.dsl.module
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.util.concurrent.TimeUnit

val API_URL1 = "http://192.168.0.6:3000"

val remoteDataSourceModule = module {
    single { createOkHttpClient() }
    single { createWebService<RestaurantsInterface>(get(), API_URL1) }
}

fun createOkHttpClient(): OkHttpClient {
    return getOkHttpClient()
}

fun getOkHttpClient(): OkHttpClient{
    return OkHttpClient
        .Builder()
        .addInterceptor(HeaderInterceptor())
        .apply {
            addInterceptor(HttpLoggingInterceptor().apply {
                level = HttpLoggingInterceptor.Level.BODY
            })
        }
        .connectTimeout(300, TimeUnit.SECONDS)
        .readTimeout(300, TimeUnit.SECONDS)
        .writeTimeout(300, TimeUnit.SECONDS)
        .build()
}

inline fun <reified T> createWebService(okHttpClient: OkHttpClient, url: String): T {
    val retrofit = Retrofit.Builder()
        .baseUrl(url)
        .addConverterFactory(GsonConverterFactory.create())
        .addCallAdapterFactory(CoroutineCallAdapterFactory())
        .client(okHttpClient)
        .build()
    return retrofit.create(T::class.java)
}