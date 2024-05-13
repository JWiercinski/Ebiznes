import './App.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import Main1, {LoginProvider, TokenProvider, UserProvider} from "./Main1";
import LoginClassic from "./LoginClassic";
import Register from "./Register";

function App() {
  return (
    <div className="App">
      <div className="Content">
          <LoginProvider>
              <TokenProvider>
                  <UserProvider>
                    <BrowserRouter>
                        <Routes>
                            <Route path="/" element={<Main1/>}></Route>
                            <Route path="/login" element={<LoginClassic/>}></Route>
                            <Route path="/register" element={<Register/>}></Route>
                        </Routes>
                    </BrowserRouter>
                  </UserProvider>
              </TokenProvider>
          </LoginProvider>
      </div>
    </div>
  );
}

export default App;
