import React from "react";
import {LoginContext, TokenContext, UserContext} from "./Main1";
import {Link, useLocation, useParams, useSearchParams} from "react-router-dom";
import axios from "axios";

function Ghidden () {
    //http://localhost:3000/ghidden/?id=1
    const { isLoggedIn, setIsLoggedIn } = React.useContext(LoginContext);
    const {token, setToken} = React.useContext(TokenContext)
    const {user, setUser} = React.useContext(UserContext)
    const [info, setInfo] = React.useState("")
    const [searchParams, setSearchParams] = useSearchParams();

    const finalize = async(event) => {event.preventDefault()
        try {
            const data = {
                ID: searchParams.get("id")
            }
            const response = await axios.post("http://localhost:22222/gdata", data)
            if (response.data.USER != null && response.data.TOKEN != null){
                setUser(response.data.USER)
                setToken(response.data.TOKEN)
                setIsLoggedIn(true)
            }
            window.location.replace("/")
        }
        catch(e){setInfo(e.response.data.message)}
    }

    return (
        <div>
            <h1>Sukces</h1>
            <button onClick={finalize}>Kontynuuj</button>
            <h2>{info}</h2>
        </div>
    )
}
export default Ghidden