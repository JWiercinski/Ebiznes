import './App.css';
import React from 'react';
import axios from "axios";
export const MessageContext = React.createContext()
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
const ChatComponent = () => {
    const { messages, addMessage } = React.useContext(MessageContext);
    const [message, setMessage] = React.useState("");
    const [info, setInfo] = React.useState(" ")

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
                    onChange={(e) => setMessage(e.target.value)}
                    placeholder="Message to the bot"
                    required
                />
                <button type="submit">Continue</button>
            </form>
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
        <div className="Chat with our llama">
        <h1>Welcome to our our LlamaBot</h1>
        <ChatComponent/>
      </div>
    </MessageProvider>
  );
}

export default App;
