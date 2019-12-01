package ni.devotion.mvvm.model


import androidx.annotation.Keep

@Keep
data class Restaurant(
    val id: Int,
    val name: String,
    val address: String,
    val phone: Int,
    val image: String
)