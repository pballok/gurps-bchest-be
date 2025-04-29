CREATE TABLE IF NOT EXISTS `character`(
    char_name VARCHAR(127) NOT NULL,
    campaign VARCHAR(127) NOT NULL,
    player VARCHAR(127) NOT NULL,
    points INT NOT NULL,
    st_modif FLOAT NOT NULL DEFAULT 0.0,
    dx_modif FLOAT NOT NULL DEFAULT 0.0,
    iq_modif FLOAT NOT NULL DEFAULT 0.0,
    ht_modif FLOAT NOT NULL DEFAULT 0.0,
    hp_modif FLOAT NOT NULL DEFAULT 0.0,
    currhp_modif FLOAT NOT NULL DEFAULT 0.0,
    will_modif FLOAT NOT NULL DEFAULT 0.0,
    per_modif FLOAT NOT NULL DEFAULT 0.0,
    fp_modif FLOAT NOT NULL DEFAULT 0.0,
    currfp_modif FLOAT NOT NULL DEFAULT 0.0,
    bs_modif FLOAT NOT NULL DEFAULT 0.0,
    bm_modif FLOAT NOT NULL DEFAULT 0.0,

    PRIMARY KEY (char_name, campaign)
)
