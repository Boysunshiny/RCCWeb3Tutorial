import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import { RustDemo2 } from "../target/types/rust_demo2";
import { BN } from "bn.js";
import { Keypair, PublicKey, SystemProgram } from "@solana/web3.js";
import { homedir } from "os";
import { join } from "path";
import { readFileSync } from "fs";

describe("rust-demo2", () => {
  // Configure the client to use the local cluster.
  anchor.setProvider(anchor.AnchorProvider.env());
  const program = anchor.workspace.RustDemo2 as Program<RustDemo2>;

  const solanaConfigPath = join(homedir(), ".config", "solana", "id.json");
  const secretKey = Uint8Array.from(
    JSON.parse(readFileSync(solanaConfigPath, "utf-8"))
  );
  const user = Keypair.fromSecretKey(secretKey);

  it("set favorites!", async () => {
    const [PDA] = PublicKey.findProgramAddressSync(
      [Buffer.from("favorites"), user.publicKey.toBuffer()],
      program.programId
    );
    console.log("user address             ", user.publicKey);
    console.log("program address          ", program.programId);
    console.log("PDA address              ", PDA);
    console.log("SystemProgram address    ", SystemProgram.programId);

    const accounts = {
      user: user.publicKey,
      favorites: PDA,
      systemProgram: SystemProgram.programId, // 注意这里是systemProgram不是system_program
    };

    console.log("accounts                 ", accounts);
    const tx = await program.methods
      .setFavorites(new BN(6), "red", ["read", "walk", "sing"])
      .accounts(accounts)
      .rpc();
    console.log("Your transaction signature", tx);
  });
});
