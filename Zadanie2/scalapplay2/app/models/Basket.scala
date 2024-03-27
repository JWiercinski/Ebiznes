package models
import models.Products
import scala.collection.mutable.ListBuffer

case class Basket(var user: String, var index: Int, var products: ListBuffer[Products])
