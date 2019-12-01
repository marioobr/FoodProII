package ni.devotion.mvvm.di

import ni.devotion.mvvm.data.network.remoteDataSourceModule
import ni.devotion.mvvm.repo.DepartmentsRepository
import ni.devotion.mvvm.repo.DepartmentsRepositoryImpl
import ni.devotion.mvvm.repo.RestaurantasRepository
import ni.devotion.mvvm.repo.RestaurantsRepositoryImpl
import ni.devotion.mvvm.viewModel.DepartmentsViewModel
import ni.devotion.mvvm.viewModel.RestaurantViewModel
import org.koin.androidx.viewmodel.dsl.viewModel
import org.koin.dsl.module

val appModule = module {
    viewModel { RestaurantViewModel(get()) }

    single<RestaurantasRepository> {
        RestaurantsRepositoryImpl(get())
    }
}

val allAppModules = listOf(appModule, remoteDataSourceModule)