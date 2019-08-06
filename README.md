# GoTest

basic http server for coupons
has methods
/coupons/ GET
/coupons/ HEAD
/coupons/filtered?[filter1]?[filter2] GET
/coupons/delete/ DELETE
/coupons/put/ handler.PUT

## PUT 
Put method should have body if the type:
{  
   "update":{  //items we want to update
      "2":{  //we need to know id of the item to update it
         "name":"Save £10 at Tesco",
         "brand":"Tesco",
         "value":10,
         "createdAt":"2018-03-01 10:15:53",
         "expiry":"2019-03-01 10:15:53"
      }
   },
   "add":[  
      {  //new item will be created and new id will be assigned
         "name":"Save £20 at Tesco",
         "brand":"Tesco",
         "value":20,
         "createdAt":"2018-03-01 10:15:53",
         "expiry":"2019-03-01 10:15:53"
      }
   ]
}

##Filter
Filter has signature: "name" eq/mt/lt "value"
* filter names are : 
* * "name" 
* * "brand"  
* * "value" 
* * "createdAt" 
* * "expiry"
* spaces should be converted to %20
* time format : "2006-01-02 15:04:05"