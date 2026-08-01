import DealChat from "./views/deal-chat/DealChat.tsx";
import BackgroundImage from "./assets/background.png";

export const App = () => <div
    className="w-full h-full"
    style={{
        backgroundPosition: "center center",
        backgroundSize: "contain",
        backgroundRepeat: "no-repeat",
        backgroundImage: `url(${BackgroundImage})`,
    }}
>
    <DealChat/>
</div>

export default App
