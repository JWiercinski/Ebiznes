import React from "react";
import {Link} from "react-router-dom";
import axios from "axios";
import {LoginContext, TokenContext, UserContext} from "./Main1";
function LoginClassic(){
 const [username, setUsername] = React.useState("")
 const [password, setPassword] = React.useState("")
    const [info, setInfo] = React.useState("")
    const {user, setUser} = React.useContext(UserContext)
    const {isLoggedIn, setIsLoggedIn} = React.useContext(LoginContext)
    const {token, setToken} = React.useContext(TokenContext)
 const sendonsubmit = async (event) =>{event.preventDefault();
  const data = {
   USER: username,
   PASSWORD: password
  }
     try {
         const response = await axios.post("http://localhost:22222/classic", data)
         if (response.data.USER != null && response.data.TOKEN != null )
         {
             setUser(response.data.USER)
             setToken(response.data.TOKEN)
             setIsLoggedIn(true)
             window.location.replace("/")
         }
         else {setInfo("Niewłaściwe dane logowania")}
     }
    catch (e){
        setInfo('Niewłaściwe dane logowania');

    }
 }

 return(
     <div className="ClassicLogin">
         <Link to="/">
             <button>Strona główna</button>
         </Link>
         <h1>Zaloguj się</h1>
         <form onSubmit={sendonsubmit}>
             <input type="text" value={username} minLength="5" maxLength="30" onChange={(e) => setUsername(e.target.value)}
                    placeholder="Nazwa użytkownika"
                    required></input>
             <input type="password" value={password} minLength="10" maxLength="50"
                    onChange={(e) => setPassword(e.target.value)} placeholder="Hasło"
                    required></input>
             <button type="submit">Zaloguj się</button>
         </form>
         <h2>{info}</h2>
     </div>
 )
}

export default LoginClassic