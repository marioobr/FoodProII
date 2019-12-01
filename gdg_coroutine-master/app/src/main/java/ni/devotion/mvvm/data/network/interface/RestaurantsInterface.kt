package ni.devotion.mvvm.data.network.`interface`

import ni.devotion.mvvm.model.Combo
import ni.devotion.mvvm.model.Department
import ni.devotion.mvvm.model.Restaurant
import retrofit2.http.GET

interface RestaurantsInterface {
    @GET("rest/obtenerRests")
    suspend fun requestRestaurants(): List<Restaurant>
    @GET("rest/obtenerCombos")
    suspend fun requestCombos(): List<Combo>
}