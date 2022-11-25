--The datees when stocks were bought
SELECT S.symbol, S.date_acquiered
FROM stocks AS S;

--Stocks where I own more than 200 shares
SELECT S.symbol, S.num_shares 
From stocks AS S
WEHRE S.num_shares > 200;