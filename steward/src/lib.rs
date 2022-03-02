//! Steward
//!
//! Application based on the [Abscissa] framework.
//!
//! [Abscissa]: https://github.com/iqlusioninc/abscissa

// Tip: Deny warnings with `RUSTFLAGS="-D warnings"` environment variable in CI

#![forbid(unsafe_code)]
#![warn(rust_2018_idioms, trivial_casts)]

pub mod allocation;
pub mod application;
pub mod cellar_aave_wrapper;
pub mod cellars;
pub mod commands;
pub mod config;
pub mod erc20;
pub mod error;
pub mod gas;
pub mod prelude;
pub mod server;
pub mod somm_send;
pub mod uniswap_pool;
pub mod uniswap_router;
pub mod utils;
