SELECT
    toStartOfMonth(date) AS month,
    sumIf(amount, amount < 0) AS expenses,
    sumIf(amount, amount > 0) AS earnings
FROM
    transactions
WHERE
    primary_class != 'Savings'
GROUP BY
    month
ORDER BY
    month;


SELECT
    date as d,
    primary_class,
    abs(sum(amount)) AS total_expenses
FROM
    transactions
WHERE
    amount < 0 AND primary_class != 'Savings' AND
    d >= toDate(parseDateTimeBestEffort('${__from:date}')) AND d <= toDate(parseDateTimeBestEffort('${__to:date}'))
GROUP BY
    d, primary_class
ORDER BY
    total_expenses DESC;



SELECT
    date as time,
    primary_class,
    abs(sum(amount)) AS total_expenses
FROM
    transactions
WHERE
    amount < 0 AND primary_class != 'Savings' AND
    time >= toDate(parseDateTimeBestEffort('${__from:date}')) AND time <= toDate(parseDateTimeBestEffort('${__to:date}'))
GROUP BY
    time, primary_class
ORDER BY
    time ASC;