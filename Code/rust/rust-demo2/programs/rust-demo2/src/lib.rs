use anchor_lang::prelude::*;

declare_id!("8tgECJCDaS747gMmRFCJkAdHZZbq5BA7tBcHmGUrrXcZ");

pub const ANCHOR_DISCRIMINATOR_SIZE: usize = 8;

#[program]
pub mod rust_demo2 {
    use super::*;
    pub fn set_favorites(
        context: Context<SetFavorites>,
        number: u64,
        color: String,
        hobbies: Vec<String>,
    ) -> Result<()> {
        msg!("欢迎   {}", context.program_id);
        let user_public_key = context.accounts.user.key();
        msg!(
            "用户 {user_public_key} ,favorite number is {number} ,favorite color is {color}, favorite hobbies is {hobbies:#?}"
        );
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
