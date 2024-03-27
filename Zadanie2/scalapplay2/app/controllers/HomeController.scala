package controllers

import models.Products
import javax.inject._
import play.api._
import play.api.mvc._
import scala.collection.mutable.ListBuffer
import scala.util.Try
import models.Category
import models.Basket

/**
 * This controller creates an `Action` to handle HTTP requests to the
 * application's home page (and a bit more, though nowhere near as clean as the template)
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
      removeitemfromelsewhere(total)
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
        removeitemfromelsewhere(xp)
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
  def removeitemfromelsewhere(index: Int): Unit = {
    val item=prodlist(index)
    for (rep <- 0 to categorieslist.size-1)
      {
        var currentC=categorieslist(rep)
        if (currentC.products.contains(item)==true)
          {
            var remid=currentC.products.indexOf(item)
            currentC.products.remove(remid)
          }
      }
    for (trial <- 0 to bucketlist.size-1)
      {
        var currentB=bucketlist(trial)
        if (currentB.products.contains(item)==true)
        {
          var remid=currentB.products.indexOf(item)
          currentB.products.remove(remid)
        }
      }
  }
  val categorieslist=new ListBuffer[Category]()
  val bucketlist=new ListBuffer[Basket]()
  categorieslist += new Category(name="Inne", products=new ListBuffer[Products](), index=0)
  bucketlist+=new Basket(user="TheOneAndOnly", products=new ListBuffer[Products](), index=0)
  categorieslist(0).products+=prodlist(0)
  bucketlist(0).products+=prodlist(0)
  val cidlist = new ListBuffer[Int]()
  cidlist += 0
  val bidlist = new ListBuffer[Int]()
  bidlist += 0
  def cgetall() = Action { implicit request =>
    Ok(views.html.cgetall(categorieslist))
  }
  def bgetall() = Action { implicit request =>
    Ok(views.html.bgetall(bucketlist))
  }
  def cidlookup(idcheck: Int): Boolean = {
    var ok=false
    if (cidlist.contains(idcheck)==true && idcheck>=0)
      ok=true
    return ok
  }
  def bidlookup(idcheck: Int): Boolean = {
    var ok=false
    if (bidlist.contains(idcheck)==true && idcheck>=0)
      ok=true
    return ok
  }
  def cgetone(cid: String) = Action {
    if (cid.forall(Character.isDigit)) {
      val pid2 = cid.toInt
      val ok = cidlookup(pid2)
      if (ok == true) {
        var id=cidlist.indexOf(pid2)
        val category= categorieslist.lift(id)
        Ok(views.html.cgetone(category))
      }
      else
        BadGateway(s"Nie udało się oðebrać kategorii o podanym indeksie")
    }
    else
      BadGateway(s"Nie udało się oðebrać kategorii o podanym indeksie")
  }
  def bgetone(pid: String) = Action {
    if (pid.forall(Character.isDigit)) {
      val pid2 = pid.toInt
      val ok = bidlookup(pid2)
      if (ok == true) {
        var id=bidlist.indexOf(pid2)
        val basket= bucketlist.lift(id)
        Ok(views.html.bgetone(basket))
      }
      else
        BadGateway(s"Nie udało się oðebrać koszyka o podanym indeksie")
    }
    else
      BadGateway(s"Nie udało się oðebrać koszyka o podanym indeksie")
  }
  def ccreate() = Action { implicit request =>
    val pro = request.body.asFormUrlEncoded
    val name = pro.get("cname").lift(0).get.toString
    var cool=true
    if (name.isEmpty==false) {
      var index=0
      if (cidlist.isEmpty==false)
      {
        index=cidlist.max + 1
      }
      for (i <- 0 to categorieslist.size-1)
      {
        if (categorieslist(i).name==name)
        {
          cool=false
        }
      }
      if (cool==true) {
        var newitem = new Category(index = index, name = name, products = new ListBuffer[Products]())
        categorieslist += newitem
        cidlist += index
        Created("Dodano element do listy")
      }
      else
        BadRequest("Istnieje już kategoria o tej nazwie")
    }
    else
      BadRequest(s"Elementu nie udało się dodać do listy... $name")
  }
  def bcreate() = Action { implicit request =>
    val pro = request.body.asFormUrlEncoded
    val name = pro.get("bname").lift(0).get.toString
    if (name.isEmpty==false) {
      var index=0
      if (bidlist.isEmpty==false)
      {
        index=bidlist.max + 1
      }
      var cool=true
      for (i <- 0 to bucketlist.size-1)
      {
        if (bucketlist(i).user==name)
        {
          cool=false
        }
      }
      if (cool==true) {
        var newitem = new Basket(index = index, user = name, products = new ListBuffer[Products]())
        bucketlist += newitem
        bidlist += index
        Created("Dodano element do listy")
      }
      else
        BadRequest("Istnieje już użytkownik o takiej nazwie")
    }
    else
      BadRequest(s"Elementu nie udało się dodać do listy... $name")
  }
  def cmod()= Action { implicit request =>
    val pro = request.body.asFormUrlEncoded
    val name = pro.get("cname").lift(0).get.toString
    val prodID = pro.get("prodid").lift(0).get.toString
    val option = pro.get("opt").lift(0).get.toString
    // option values - r to remove product from category, a to add
    val index2 = pro.get("cmid").lift(0).get.toString
    var ok=false
    if (index2.isEmpty==false && index2.forall(Character.isDigit)==true) {
      val index = index2.toInt
      ok = cidlookup(index)
    }
    if (ok==true) {
      val index=index2.toInt
      var newindex = cidlist.indexOf(index)
      if (name.isEmpty==false) {
        categorieslist(newindex).name = name
      }
      if (prodID.isEmpty==false && option.isEmpty==false) {
        if (prodID.forall(Character.isDigit)==true)
          {
            val productID=prodID.toInt
            if (pidlookup(productID)==true)
              {
                var pid=pidlist.indexOf(productID)
                option match {
                  case "r" =>{
                    if (categorieslist(newindex).products.contains(prodlist(pid))==true){
                  categorieslist(newindex).products.remove(pid)
                  }}
                  case "a" =>{
                    if (categorieslist(newindex).products.contains(prodlist(pid))==false)
                      {
                  categorieslist(newindex).products+=prodlist(pid)
                  }}
                }
            }
          }
        }
      val category = categorieslist.lift(newindex)
      Ok(views.html.cgetone(category))
    }
    else
    {
      NotFound
    }
  }
  def crem(cid: String)= Action {
    if (cid.forall(Character.isDigit)) {
      val cid2 = cid.toInt
      val ok = cidlookup(cid2)
      if (ok == true) {
        var xp = cidlist.indexOf(cid2)
        categorieslist.remove(xp)
        cidlist-=xp
        NoContent
      }
      else
        NotFound(s"Nie udało się usunąć kategorii o podanym indeksie")
    }
    else
      NotFound(s"Nie udało się usunąć kategorii o podanym indeksie")
  }
  def brem(bid: String)= Action {
    if (bid.forall(Character.isDigit)) {
      val bid2 = bid.toInt
      val ok = bidlookup(bid2)
      if (ok == true) {
        var xp = bidlist.indexOf(bid2)
        bucketlist.remove(xp)
        bidlist-=xp
        NoContent
      }
      else
        NotFound(s"Nie udało się usunąć koszyka o podanym indeksie")
    }
    else
      NotFound(s"Nie udało się usunąć koszyka o podanym indeksie")
  }
  def bmod()= Action { implicit request =>
    val pro = request.body.asFormUrlEncoded
    val name = pro.get("user").lift(0).get.toString
    val prodID = pro.get("prodid").lift(0).get.toString
    val option = pro.get("opt").lift(0).get.toString
    // option values - r to remove product from category, a to add
    val index2 = pro.get("bmid").lift(0).get.toString
    var ok = false
    if (index2.isEmpty == false && index2.forall(Character.isDigit) == true) {
      val index = index2.toInt
      ok = bidlookup(index)
    }
    if (ok == true) {
      val index = index2.toInt
      var newindex = bidlist.indexOf(index)
      if (name.isEmpty == false) {
        bucketlist(newindex).user = name
      }
      if (prodID.isEmpty == false && option.isEmpty == false) {
        if (prodID.forall(Character.isDigit) == true) {
          val productID = prodID.toInt
          if (pidlookup(productID) == true) {
            var pid = pidlist.indexOf(productID)
            option match {
              case "r" => {
                if (bucketlist(newindex).products.contains(prodlist(pid)) == true) {
                  bucketlist(newindex).products.remove(pid)
                }
              }
              case "a" => {
                bucketlist(newindex).products += prodlist(pid)
              }
            }
          }
        }
      }
        val basket = bucketlist.lift(newindex)
        Ok(views.html.bgetone(basket))
      }
      else {
        NotFound
      }
    }
}