// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateToken } from "./types/token/tx";
import { MsgUpdateToken } from "./types/token/tx";
import { MsgDeleteToken } from "./types/token/tx";


const types = [
  ["/Palo_Alt0.vanillaneyx.token.MsgCreateToken", MsgCreateToken],
  ["/Palo_Alt0.vanillaneyx.token.MsgUpdateToken", MsgUpdateToken],
  ["/Palo_Alt0.vanillaneyx.token.MsgDeleteToken", MsgDeleteToken],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateToken: (data: MsgCreateToken): EncodeObject => ({ typeUrl: "/Palo_Alt0.vanillaneyx.token.MsgCreateToken", value: MsgCreateToken.fromPartial( data ) }),
    msgUpdateToken: (data: MsgUpdateToken): EncodeObject => ({ typeUrl: "/Palo_Alt0.vanillaneyx.token.MsgUpdateToken", value: MsgUpdateToken.fromPartial( data ) }),
    msgDeleteToken: (data: MsgDeleteToken): EncodeObject => ({ typeUrl: "/Palo_Alt0.vanillaneyx.token.MsgDeleteToken", value: MsgDeleteToken.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
