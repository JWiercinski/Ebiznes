import './App.css';
import React from 'react';
import axios from "axios";
export const MessageContext = React.createContext()
export const StartedContext = React.createContext()
export const MessageProvider = ({ children }) => {
    const [messages, setMessages] = React.useState(() => {
        const storedMessages = JSON.parse(localStorage.getItem("messages")) || [];
        return storedMessages;
    });

    React.useEffect(() => {
        localStorage.setItem("messages", JSON.stringify(messages));
    }, [messages]);

    const addMessage = (newMessage) => {
        setMessages((prevMessages) => [...prevMessages, newMessage]);
    };

    return (
        <MessageContext.Provider value={{ messages, addMessage }}>
            {children}
        </MessageContext.Provider>
    );
};
export const StartedProvider = ({children}) =>
{
    const getInitialState = () => {
        const savedState = localStorage.getItem('started');
        return savedState !== null ? JSON.parse(savedState) : false;
    };
    const [started, setStarted] = React.useState(getInitialState)

    React.useEffect(() => {
        localStorage.setItem('started', JSON.stringify(started));
    }, [started]);

    return (
        <StartedContext.Provider value={{ started, setStarted }}>
            {children}
        </StartedContext.Provider>
    );
}
const ChatComponent = () => {
    const { messages, addMessage } = React.useContext(MessageContext);
    const [message, setMessage] = React.useState("");
    const {started, setStarted} = React.useContext(StartedContext)

    const pushMessage = async (event) =>
    {
        event.preventDefault();
        try
        {
            const data = {message: message}
            const newMess = { user: "User", content: data.message };
            addMessage(newMess);
            const response = await axios.post("http://localhost:8000/chat/", data)
            addMessage({user: "LLAMA", content: response.data.response})
            window.location.reload()
        }
        catch (e)
        {
            //"Error encountered. Apologies for the inconvenience"
            addMessage({user: "SYSTEM", content: e.message})
            window.location.reload()
        }
    };

    const startChat = async (event) =>
    {
        try
        {
            const response = await axios.get("http://localhost:8000/start/")
            addMessage({user: "LLAMA", content: response.data.response})
            setStarted(true)
        }
        catch (e)
        {
            addMessage({user: "SYSTEM", content: e.message})
            window.location.reload()
        }
    }

    const endChat = async (event) =>
    {
        try
        {
            const response = await axios.get("http://localhost:8000/end/")
            addMessage({user: "LLAMA", content: response.data.response})
            setStarted(false)
        }
        catch (e)
        {
            addMessage({user: "SYSTEM", content: e.message})
            window.location.reload()
        }
    }

    const clearChat = () =>
    {
        localStorage.removeItem("messages");
        window.location.reload()
    }

    return (
        <div>
            <form onSubmit={pushMessage}>
                <input
                    type="text"
                    value={message}
                    disabled={!started}
                    onChange={(e) => setMessage(e.target.value)}
                    placeholder="Message to the bot"
                    required
                />
                <button type="submit" disabled={!started}>Continue</button>
            </form>
            <p></p>
            <button onClick={startChat} disabled={started}>Start Chat</button>
            <button onClick={endChat} disabled={!started}>End Chat</button>
            <p></p>
            <button onClick={clearChat}>Clear Chat</button>
            <p></p>
            <h2>Chat History</h2>
            {messages.map((msg, index) => (
                <p key={index}>{msg.user}: {msg.content}</p>
            ))}
        </div>
    );
};

function App() {
    return (
        <MessageProvider>
            <StartedProvider>
                <div className="Chat with our llama">
                    <h1>Welcome to our our LlamaBot</h1>
                    <ChatComponent/>
                </div>
            </StartedProvider>
        </MessageProvider>
  );
}

export default App;
