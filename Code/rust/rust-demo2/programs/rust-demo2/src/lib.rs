use anchor_lang::prelude::*;

declare_id!("GvBrcFw27V61tUgGrvr9tnp1Ajbt3ETofsm949MkUBFR");

pub const ANCHOR_DISCRIMINATOR_SIZE: usize = 8;

#[program]
pub mod rust_demo2 {
    // use futures::executor::block_on;

    use super::*;
    pub fn set_favorites(
        context: Context<SetFavorites>,
        number: u64,
        color: String,
        hobbies: Vec<String>,
    ) -> Result<()> {
        // let b = async {
        //     msg!("这是我的测试");
        // };
        // block_on(b);

        msg!("欢迎   {}", context.program_id);
        let user_public_key = context.accounts.user.key();
        // msg!(
        //     "用户 {user_public_key} ,favorite number is {number} ,favorite color is {color}, favorite hobbies is {hobbies:#?}"
        // );
        msg!(
            "用户 {} ,favorite number is {} ,favorite color is {}, favorite hobbies is {:#?}",
            user_public_key,
            number,
            color,
            hobbies
        );
        context.accounts.favorites.set_inner(Favorites {
            number,
            color,
            hobbies,
        });
        Ok(())
    }

    pub fn set_number(context: Context<SetFavorites>, number: u64) -> Result<()> {
        msg!("欢迎 {}", context.program_id);
        msg!("你爱的   {}", number);
        Ok(())
    }
    pub fn set_number_and_color(
        context: Context<SetFavorites>,
        number: u64,
        color: String,
    ) -> Result<()> {
        msg!("欢迎 {}", context.program_id);
        msg!("你爱的   {}  喜欢的颜色 {}", number, color);
        Ok(())
    }
    pub fn set_color(context: Context<SetFavorites>, color: String) -> Result<()> {
        msg!("欢迎 {}", context.program_id);
        msg!("你爱的   喜欢的颜色 {}", color);
        Ok(())
    }
}

#[account]
#[derive(InitSpace)]
pub struct Favorites {
    pub number: u64,
    #[max_len(50)]
    pub color: String,
    #[max_len(5, 50)]
    pub hobbies: Vec<String>,
}

#[derive(Accounts)]
pub struct SetFavorites<'info> {
    #[account(mut)]
    pub user: Signer<'info>,
    /*
    init_if_needed,
    payer = user,
    space = ANCHOR_DISCRIMINATOR_SIZE + Favorites::INIT_SPACE,
    seeds = [b"favorites",user.key().as_ref()],
    bump
    */
    #[account(
        init_if_needed,
        payer = user,
        space = ANCHOR_DISCRIMINATOR_SIZE + Favorites::INIT_SPACE,
        seeds = [b"favorites",user.key().as_ref()],
        bump
        )]
    pub favorites: Account<'info, Favorites>,
    pub system_program: Program<'info, System>,
}
