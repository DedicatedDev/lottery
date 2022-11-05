import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBidToLottery } from "./types/lottery/lottery/tx";
import { MsgCreateLottery } from "./types/lottery/lottery/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dedicateddev.lottery.lottery.MsgBidToLottery", MsgBidToLottery],
    ["/dedicateddev.lottery.lottery.MsgCreateLottery", MsgCreateLottery],
    
];

export { msgTypes }