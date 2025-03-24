use anchor_lang::prelude::*;

declare_id!("GvBrcFw27V61tUgGrvr9tnp1Ajbt3ETofsm949MkUBFR");

#[program]
pub mod rust_demo2 {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        msg!("Greetings from: {:?}", ctx.program_id);
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize {}
