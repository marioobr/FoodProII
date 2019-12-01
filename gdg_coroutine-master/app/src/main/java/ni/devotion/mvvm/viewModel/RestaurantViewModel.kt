package ni.devotion.mvvm.viewModel

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.launch
import ni.devotion.mvvm.R
import ni.devotion.mvvm.model.Department
import ni.devotion.mvvm.model.Restaurant
import ni.devotion.mvvm.repo.DepartmentsRepository
import ni.devotion.mvvm.repo.RestaurantasRepository

class RestaurantViewModel (private val restaurantasRepository: RestaurantasRepository): ViewModel() {
    private val _uiState = MutableLiveData<BusinessDataState>()
    val uiState: LiveData<BusinessDataState> get() = _uiState

    init {
        retrieveBusiness()
    }

    fun retrieveBusiness(){
        viewModelScope.launch {
            runCatching {
                emitUiState(showProgress = true)
                restaurantasRepository.getRestaurants()
            }.onSuccess {
                emitUiState(rest = Event(it))
            }.onFailure {
                emitUiState(error = Event(R.string.internet_failure_error))
            }
        }
    }

    private fun emitUiState(showProgress: Boolean = false, rest: Event<List<Restaurant>>? = null, error: Event<Int>? = null){
        val dataState = BusinessDataState(showProgress, rest, error)
        _uiState.value = dataState
    }

    data class BusinessDataState(val showProgress: Boolean, val restaurant: Event<List<Restaurant>>?, val error: Event<Int>?)
}