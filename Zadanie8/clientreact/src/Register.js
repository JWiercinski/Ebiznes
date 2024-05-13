import React from "react";
import {Link} from "react-router-dom";
import axios from "axios";
import {LoginContext, TokenContext, UserContext} from "./Main1";
function Register(){
    const [username, setUsername] = React.useState("")
    const [name, setName] = React.useState("")
    const [password, setPassword] = React.useState("")
    const [email, setEmail] = React.useState("")
    const [info, setInfo] = React.useState("")
    const {user, setUser} = React.useContext(UserContext)
    const {isLoggedIn, setIsLoggedIn} = React.useContext(LoginContext)
    const {token, setToken} = React.useContext(TokenContext)
    const sendonsubmit = async (event) =>{event.preventDefault();
        const data = {
            USER: username,
            NAME: name,
            EMAIL: email,
            PASSWORD: password
        }
        try {
            const response = await axios.post("http://localhost:22222/register", data)
            if (response.data.USER != null && response.data.TOKEN != null )
            {
                setUser(response.data.USER)
                setToken(response.data.TOKEN)
                setIsLoggedIn(true)
                window.location.replace("/")
            }
            else{
                setInfo("Nastąpił nieznany błąd podczas próby rejestracji")
            }
        }
        catch (e){
            setInfo(e.response.data);
        }
    }

    return(
        <div className="Register">
            <Link to="/">
                <button>Strona główna</button>
            </Link>
            <h1>Zarejestruj się</h1>
            <form onSubmit={sendonsubmit}>
                <input type="text" value={username} minLength="2" maxLength="30" onChange={(e) => setUsername(e.target.value)}
                       placeholder="Nazwa użytkownika"
                       required></input>
                <input type="text" value={name} minLength="3" maxLength="30" onChange={(e) => setName(e.target.value)}
                       placeholder="Imię i nazwisko"
                       required></input>
                <input type="text" value={email} minLength="7" maxLength="50"
                       onChange={(e) => setEmail(e.target.value)} placeholder="E-mail"
                       required></input>
                <input type="password" value={password} minLength="10" maxLength="50"
                       onChange={(e) => setPassword(e.target.value)} placeholder="Hasło"
                       required></input>
                <button type="submit">Zarejestruj się</button>
            </form>
            <h2>{info}</h2>
        </div>
    )
}

export default Register