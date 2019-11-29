## Project: Infinity
## File: model design

#### User model description 

````{
Uuid: unique identifier
Email: user email address
FirstName: 
LastName:
PhoneNumber:
Address: {
	Lat:	
	Lng:
	City:
	State:
	Country:
	Pin:
	}
Type: user type (buyer, seller, admin)
Password: ****
CreatedAt: when was the profile created
UpdatedAt: when was the profile last updated
}
````

##### Functions on User Model

1. Singup
2. Login
3. Email verification
4. Reset password
5. Forgot password

-----------


#### Item model description: 
````{
Uuid: unique identifier
Id: customer id easy to remember and search eg. (Item1001, Item1002)
Type: type of item
Organic: true or false
Name: Name of the Item
CreatedBy: userId ho created the Item
Description: Item description
}
````

##### Functions on Item model

1. Create item
2. Update item
3. Remove item

-----

#### MarketPlace model description:

````{
ItemId:
Quantity:
ExpiryDate:
Description:
pricePerUnit:
Unit: [Kg, gram etc]
Address: {
	Lat:	
	Lng:
	City:
	State:
	Country:
	Pin:
	}
CreatedAt: 
UpdatedAt:
}
````


##### Functions on MarketPlace model

1. Add Item for selling
2. Update Item for selling (to change quantity or description or price)
3. Decrease quantity after sold
4. Remove from market place

