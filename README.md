# Where it's deployed
https://dmitryrn.github.io/order_creator/

# Application description
Imagine for a moment that one of our product lines ships in various pack sizes:
• 250 Items
• 500 Items
• 1000 Items
• 2000 Items
• 5000 Items
Our customers can order any number of these items through our website, but they will always only
be given complete packs.
1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each
order.
(Please note, rule #2 takes precedence over rule #3)

| **Items ordered** | **Correct number of packs** | **Incorrect number of packs** |
|------------------|--------------------------|------------------------------|
| 1 | 1 × 250 | 1 × 500 – more items than necessary |
| 250 | 1 × 250 | 1 × 500 – more items than necessary |
| 251 | 1 × 500 | 2 × 250 – more packs than necessary |
| 501 | 1 × 500 <br> 1 × 250 | 1 × 1000 – more items than necessary <br> 3 × 250 – more packs than necessary |
| 12001 | 2 × 5000 <br> 1 × 2000 <br> 1 × 250 | 3 × 5000 – more items than necessary |

*(Please note, rule #2 takes precedence over rule #3.)*

# Other
Template used: https://github.com/gkatanacio/go-serverless-template
