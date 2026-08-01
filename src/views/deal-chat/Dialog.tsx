import Avatar from "../../components/Avatar.tsx";
import ChatItem from "./ChatItem.tsx";

import GoblinNormal from "../../assets/goblin/avatar-000.png"

export const Dialog = (props: {
    chatItems: object[]  // TODO: Define Chat Item Interface
}) => {
    const {chatItems} = props
    {/*<div className="rpg-window h-[50vh] min-h-64 my-6 mx-16 p-8">*/}
    return <div className="rpg-window felx min-h-0 flex-col overflow-hidden my-6 mx-16 p-8">
        <div
            className="
            min-h-42
            flex items-center justify-between border-b-2
            border-[rgba(232,217,166,0.35)] p-4 mb-4 font-bold text-cream-500 uppercase">
            <span>Selling NFT to Mob Goblin</span>
            <Avatar src={GoblinNormal}/>
        </div>
        <div className="rpg-log gap-3 max-h-[35vh] min-h-60">
            {chatItems.map((chatItem, i) => (
                <div key={i} className="my-2"><ChatItem chatItem={chatItem}/></div>
            ))}
        </div>
    </div>
}

export default Dialog