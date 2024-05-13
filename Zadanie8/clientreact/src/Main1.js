import React, { useState } from 'react';
import {Link} from "react-router-dom";

export const LoginContext = React.createContext();
export const TokenContext = React.createContext();
export const UserContext = React.createContext();
export const LoginProvider = ({ children }) => {
    const getInitialLoginState = () => {
        const savedLoginState = localStorage.getItem('isLoggedIn');
        return savedLoginState !== null ? JSON.parse(savedLoginState) : false;
    };

    const [isLoggedIn, setIsLoggedIn] = useState(getInitialLoginState);

    React.useEffect(() => {
        localStorage.setItem('isLoggedIn', JSON.stringify(isLoggedIn));
    }, [isLoggedIn]);

    return (
        <LoginContext.Provider value={{ isLoggedIn, setIsLoggedIn }}>
            {children}
        </LoginContext.Provider>
    );
};
export const TokenProvider = ({ children }) => {
    const getInitialTokenS = () => {
        const savedTokenState = localStorage.getItem('token');
        return savedTokenState !== null ? JSON.parse(savedTokenState) : null;
    };

    const [token, setToken] = useState(getInitialTokenS);

    React.useEffect(() => {
        localStorage.setItem('token', JSON.stringify(token));
    }, [token]);

    return (
        <TokenContext.Provider value={{ token, setToken }}>
            {children}
        </TokenContext.Provider>
    );
};
export const UserProvider = ({ children }) => {
    const getInitialUser = () => {
        const savedUser = localStorage.getItem('user');
        return savedUser !== null ? JSON.parse(savedUser) : null;
    };

    const [user, setUser] = useState(getInitialUser);

    React.useEffect(() => {
        localStorage.setItem('user', JSON.stringify(user));
    }, [user]);

    return (
        <UserContext.Provider value={{ user, setUser}}>
            {children}
        </UserContext.Provider>
    );
};
function Main1() {
    const { isLoggedIn, setIsLoggedIn } = React.useContext(LoginContext);
    const {token, setToken} = React.useContext(TokenContext)
    const {user, setUser} = React.useContext(UserContext)

    if (token === null && setIsLoggedIn === true)
    {
        setIsLoggedIn(false)
        setUser(null)
        window.location.reload()
    }

    const handleLogout = () => {
        setIsLoggedIn(false);
        setToken(null)
        setUser(null)
    };

    return (
        <div>
            <h1>Witamy {user}</h1>
            <p>Na tej stronie możesz się zalogować lub wylogować</p>
            <Link to="/login"><button disabled={isLoggedIn}>Zaloguj się</button></Link>
            <Link to="/register"><button disabled={isLoggedIn}>Zarejestruj się</button></Link>
            <button onClick={handleLogout} disabled={!isLoggedIn}>Wyloguj się</button>
        </div>
    );
}
export default Main1;