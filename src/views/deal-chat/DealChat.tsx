import Dialog from "./Dialog.tsx";
import Avatar from "../../components/Avatar.tsx";


import MerchantNormal from "../../assets/merchant/avatar-200.png"
import PriestNormal from "../../assets/priest/avatar-020.png"
import MageNormal from "../../assets/mage/avatar-009.png"


export const DealChat = () => {
    return <div
        className="w-full columns-1 gap-24 backdrop-blur-sm">
        <Dialog chatItems={[{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}]}/>
        {/* Input Area */}
        <div className="w-full relative bg-forest-800 mt-18 z-0">
            <div className="w-full absolute top-8">
                <div className="w-full flex flex-row absolute -top-24 gap-2">
                    <div className="mx-8 my-2 -z-10 flex-1 text-left  justify-items-start"><Avatar src={PriestNormal}/>
                    </div>
                    <div className="mx-8 my-2 -z-10 flex-1 text-right justify-items-end"><Avatar src={MageNormal}/>
                    </div>
                </div>
                <div className="rpg-window flex flex-row h-fit my-6 mx-16 p-8 gap-2">
                    <div className="flex flex-col flex-1 align-middle min-w-32"><Avatar src={MerchantNormal}/></div>
                    <div className="flex-4 columns-1 max-w-[80%]">
                        <div
                            className="
                                text-cream-700 border-b-2 border-[rgba(232,217,166,0.35)] p-1 mb-4">
                            <span>Make an offer</span>
                        </div>
                        <div className="flex flex-row w-full gap-3 mb-4">
                            <input
                                className="rpg-input flex-9"
                                placeholder="Speak your words to the Warden..."
                            />
                            <button className="rpg-button flex-1">Send</button>
                            <button className="rpg-button flex-1">Deal</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
}

export default DealChat