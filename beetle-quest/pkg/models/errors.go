package models

type constError string

func (e constError) Error() string {
	return string(e)
}

const (
	ErrInternalServerError             constError = "internal server error"
	ErrInvalidUserID                   constError = "invalid user id"
	ErrInvalidAuctionID                constError = "invalid auction id"
	ErrCouldNotGenerateAuction         constError = "could not generate auction"
	ErrUserNotFound                    constError = "user not found"
	ErrInvalidUsernameOrPass           constError = "invalid username or password"
	ErrInvalidUsernameOrPassOrEmail    constError = "invalid username or password or email"
	ErrCheckingPassword                constError = "error when checking password"
	ErrUserParametersNotValid          constError = "inserted username or mail are already in the system"
	ErrCouldNotDelete                  constError = "could not delete user"
	ErrInvalidPassword                 constError = "invalid password"
	ErrCouldNotUseNewPassword          constError = "could not use new password"
	ErrCouldNotUpdate                  constError = "could not update user"
	ErrUsernameAlreadyExists           constError = "username already exists"
	ErrEmailAlreadyExists              constError = "email already exists"
	ErrAmountNotValid                  constError = "amount not valid"
	ErrNotEnoughMoneyToBuyGacha        constError = "not enough money to buy gacha"
	ErrInvalidEndTime                  constError = "invalid end time"
	ErrCouldNotCreateAuction           constError = "could not create auction"
	ErrRetrievingAuctions              constError = "could not retrieve auctions"
	ErrBidsNotFound                    constError = "bids not found"
	ErrUserAlreadyHasGacha             constError = "user already has gacha"
	ErrCouldNotRollGacha               constError = "could not roll gacha"
	ErrNotEnoughMoneyToRollGacha       constError = "not enough money to roll gacha"
	ErrCouldNotRetrieveUserAuctions    constError = "could not retrieve user auctions"
	ErrGachaAlreadyAuctioned           constError = "gacha already auctioned"
	ErrUserNotOwnerOfAuction           constError = "user not owner of auction"
	ErrCouldNotDeleteAuction           constError = "could not delete auction"
	ErrCouldNotRetrieveAuctionBids     constError = "could not retrieve auction bids"
	ErrAuctionHasBids                  constError = "auction has bids"
	ErrAuctionEnded                    constError = "auction already ended"
	ErrAuctionIsTooCloseToEnd          constError = "auction is too close to end"
	ErrOwnerCannotBid                  constError = "owner cannot bid"
	ErrNotEnoughMoneyToBid             constError = "not enough money to bid"
	ErrInvalidBidAmount                constError = "invalid bid amount"
	ErrCouldNotBidToAuction            constError = "could not bid to auction"
	ErrBidAmountNotEnough              constError = "bid amount not enough"
	ErrCouldNotSetCloseEventForAuction constError = "could not set close event for auction"
	ErrCouldNotUpdateAuction           constError = "could not update auction"
	ErrCouldNotAddTransaction          constError = "could not add transaction"
	ErrCouldNotAddEvent                constError = "could not add event"
	ErrCouldNotAddGachaToUser          constError = "could not add gacha to user"
	ErrCouldNotRemoveGachaFromUser     constError = "could not remove gacha from user"
	ErrTransactionNotFound             constError = "transaction not found"
	ErrMaxMoneyExceeded                constError = "max money exceeded"
	ErrInvalidAdminIDOrPassOrOTOP      constError = "invalid admin id or password or one time password"
	ErrGachaCreationFailed             constError = "gacha creation failed"
	ErrCouldNotCreateGacha             constError = "could not create gacha"
	ErrCouldNotUpdateGacha             constError = "could not update gacha"
	ErrCouldNotDeleteGacha             constError = "could not delete gacha"
	ErrCouldNotRetrieveAuctions        constError = "could not retrieve auctions"
	ErrCouldNotRetrieveTransactions    constError = "could not retrieve transactions"

	ErrUsernameOrEmailAlreadyExists constError = "username or email already exists"

	ErrGachaAlreadyExists          constError = "gacha already exists"
	ErrRetalationGachaUserNotFound constError = "retalation gacha user not found"
	ErrAdminNotFound               constError = "admin not found"
	ErrUserTransactionNotFound     constError = "user transaction not found"

	ErrCouldNotParseTime          constError = "time format not correct"
	ErrCouldNotDecodeUserID       constError = "could not decode user id"
	ErrCouldNotDecodeGachaID      constError = "could not decode gacha id"
	ErrCouldNotDecodeAuctionID    constError = "could not decode auction id"
	ErrCouldNotFindResourceByUUID constError = "could not find resource by uuid"
	ErrCouldNotRetrieveUserGachas constError = "could not retrieve user gachas"
	ErrUserDoesNotOwnGacha        constError = "user does not own gacha"
	ErrInvalidData                constError = "invalid data"
	ErrInvalidTimeFormat          constError = "invalid time format"
	ErrFailedToRevokeToken        constError = "failed to revoke token"
	ErrFailedToVerifyToken        constError = "failed to verify token"

	ErrInvalidGachaID constError = "invalid gacha id"
	ErrGachaNotFound  constError = "gacha not found"

	ErrAuctionNotFound       constError = "auction not found"
	ErrAuctionAltreadyExists constError = "auction already exists"

	// Internal api errors
	ErrCouldNotBuyGacha constError = "internal: could not buy gacha (db)"

	ErrInvalidIntValueAsString constError = "invalid int value as string"
	ErrInvalidRarityValue      constError = "invalid rarity value [Common, Uncommon, Rare, Epic, Legendary]"
)