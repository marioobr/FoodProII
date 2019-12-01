package ni.devotion.mvvm.repo

import ni.devotion.mvvm.data.network.`interface`.DepartmentsInterface
import ni.devotion.mvvm.data.network.`interface`.RestaurantsInterface
import ni.devotion.mvvm.model.Department
import ni.devotion.mvvm.model.Restaurant

interface RestaurantasRepository {
    suspend fun getRestaurants(): List<Restaurant>
}

class RestaurantsRepositoryImpl(private val restaurantService: RestaurantsInterface):
    RestaurantasRepository {
    override suspend fun getRestaurants(): List<Restaurant> {
        return restaurantService.requestRestaurants()
    }
}