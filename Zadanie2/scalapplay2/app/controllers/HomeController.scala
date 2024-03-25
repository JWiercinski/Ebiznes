package controllers
import models.Products
import javax.inject._
import play.api._
import play.api.mvc._
import scala.collection.mutable.ListBuffer
import scala.util.Try

/**
 * This controller creates an `Action` to handle HTTP requests to the
 * application's home page.
 * Component configuration - SBT 1.9.2, Scala 2.12.12, IBM JDK 1.8.0_402
 */
@Singleton
class HomeController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  /**
   * Create an Action to render an HTML page.
   *
   * The configuration in the `routes` file means that this method
   * will be called when the application receives a `GET` request with
   * a path of `/`.
   */
  def index() = Action { implicit request: Request[AnyContent] =>
    Ok(views.html.index())
  }

  val prodlist = new ListBuffer[Products]()
  prodlist += new Products(pid=0, pname="VeryDumbName", pprice=12, pman="Debvils")
  val pidlist = new ListBuffer[Int]()
  pidlist += 0

  /* Deprecated in favour of pidlookup
  def listsize(idcheck: Int): Boolean = {
    var sizer = prodlist.length
    var ok = false
    if (sizer > idcheck && idcheck >= 0) {
      ok = true
    }
    return ok
  }
  */
  def pidlookup(idcheck: Int): Boolean = {
    var ok=false
    if (pidlist.contains(idcheck)==true && idcheck>=0)
      ok=true
    return ok
  }
  def pgetall() = Action { implicit request =>
    Ok(views.html.pgetall(prodlist))
  }
  def pgetone(pid: String) = Action {
    if (pid.forall(Character.isDigit)) {
      val pid2 = pid.toInt
      val ok = pidlookup(pid2)
      if (ok == true) {
        var id=pidlist.indexOf(pid2)
        val product = prodlist.lift(id)
        Ok(views.html.pgetone(product))
      }
      else
        BadGateway(s"Nie udało się oðebrać produktu o podanym indeksie")
    }
    else
      BadGateway(s"Nie udało się oðebrać produktu o podanym indeksie")
  }

  def premoval() = Action { implicit request =>
    val index = request.getQueryString("prid").flatMap(s => Try(s.toInt).toOption)
    var ok=false
    if (index.isDefined) {
      ok = pidlookup(index.get)
    }
    if (ok == true) {
      val nice=index.get
      val total= pidlist.indexOf(nice)
      prodlist.remove(total)
      pidlist-=nice
      Ok(s"Element $nice usunięty")
    }
    else
      BadGateway(s"Nie udało się usunąć elementu o podanym indeksie")
  }

  def isprice (string: String): Boolean = {
    val pricepattern = """^\d+(\.\d{1,2})?$""".r
    string match {
      case pricepattern(_*) => true
      case _ => false
    }
  }
  def pcreate() = Action { implicit request =>
    val pro = request.body.asFormUrlEncoded
    val name = pro.get("pname").lift(0).get.toString
    val price = pro.get("pprice").lift(0).get.toString
    val manuf = pro.get("pman").lift(0).get.toString
    if (name.isEmpty==false && price.isEmpty==false && isprice(price)==true && manuf.isEmpty==false) {
      val price2=price.toDouble
      var index=0
      if (pidlist.isEmpty==false)
        {
          index=pidlist.max + 1
        }
      var newitem = new Products(pid=index, pname = name, pprice = price2, pman = manuf)
      prodlist += newitem
      pidlist += index
      Created("Dodano element do listy")
    }
    else
      BadRequest(s"Elementu nie udało się dodać do listy... $name $price $manuf")
  }
  def pmodify() = Action { implicit request =>
    val name = request.getQueryString("pname").flatMap(s => Try(s.toString).toOption)
    val price = request.getQueryString("pprice").flatMap(s => Try(s.toDouble).toOption)
    val manuf = request.getQueryString("pman").flatMap(s => Try(s.toString).toOption)
    val index = request.getQueryString("pmid").flatMap(s => Try(s.toInt).toOption)
    var ok=false
    if (index.isDefined) {
      ok = pidlookup(index.get)
    }
    if (ok==true)
      {
        var newindex = pidlist.indexOf(index.get)
        if (name.exists(_.trim.nonEmpty)) {
          prodlist(newindex).pname=name.get
        }
        if (price.isDefined)
          {
            prodlist(newindex).pprice=price.get
          }
        if (manuf.exists(_.trim.nonEmpty))
          {
            prodlist(newindex).pman=manuf.get
          }
        val product = prodlist.lift(newindex)
        Ok(views.html.pgetone(product))
      }
    else
      BadGateway(s"Elementu nie udało się zmodyfikować")
  }
  def prem(pid: String)= Action {
    if (pid.forall(Character.isDigit)) {
      val pid2 = pid.toInt
      val ok = pidlookup(pid2)
      if (ok == true) {
        var xp = pidlist.indexOf(pid2)
        val product = prodlist.lift(xp)
        prodlist.remove(xp)
        pidlist-=xp
        NoContent
      }
      else
        NotFound(s"Nie udało się usunąć produktu o podanym indeksie")
    }
    else
      NotFound(s"Nie udało się usunąć produktu o podanym indeksie")
  }
  def pmod()= Action { implicit request =>
    val pro = request.body.asFormUrlEncoded
    val name = pro.get("pname").lift(0).get.toString
    val price = pro.get("pprice").lift(0).get.toString
    val manuf = pro.get("pman").lift(0).get.toString
    val index2 = pro.get("pmid").lift(0).get.toString
    var ok=false
    if (index2.isEmpty==false && index2.forall(Character.isDigit)==true) {
      val index = index2.toInt
      ok = pidlookup(index)
    }
    if (ok==true) {
      val index=index2.toInt
      var newindex = pidlist.indexOf(index)
      if (name.isEmpty==false) {
        prodlist(newindex).pname = name
      }
      if (price.isEmpty==false) {
        if (isprice(price)==true) {
          val price2 = price.toDouble
          prodlist(newindex).pprice = price2
        }
      }
      if (manuf.isEmpty==false) {
        prodlist(newindex).pman = manuf
      }
      val product = prodlist.lift(newindex)
      Ok(views.html.pgetone(product))
    }
    else
      {
        NotFound
      }
  }
}