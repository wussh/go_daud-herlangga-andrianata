use test //membuat database
db.createCollection ('name') //membuat collection
db.products.insertMany([{
    _id:"type1",
    operator_id:3,
    barang:[{
        _id: "p1",
        name: "Indomie Goreng",
        description: "Mie Goreng",
        price: new NumberLong(2000)
    },
    {
        _id: "p2",
        name: "Mie Sedap",
        description: "Mie Goreng",
        price: new NumberLong(2000)
    }]},
    {
        _id:"type2",
        operator_id:1,
        barang:[{
            _id: "p3",
            name: "Indomie Ayam Bawang",
            description: "Mie Kuah",
            price: new NumberLong(2000)
        },
        {
            _id: "p4",
            name: "Indomie Soto",
            description: "Mie Kuah",
            price: new NumberLong(2000)
        },
        {
            _id: "p5",
            name: "Indomie Kare",
            description: "Kuah",
            price: new NumberLong(2000)
        }]},
        {
            _id:"type3",
            operator_id:4,
            barang:[{
                _id: "p6",
                name: "Ramen Pedas Lvl 1",
                description: "Mie Ramen",
                price: new NumberLong(2000)
            },
            {
                _id: "p7",
                name: "Ramen Pedas Lvl 2",
                description: "Mie Ramen",
                price: new NumberLong(2000)
            },
            {
                _id: "p8",
                name: "Ramen Pedas Lvl 3",
                description: "Mie Ramen",
                price: new NumberLong(2000)
            }]}
  
]);

db.user.insertMany([{
    _id: "u1",
    name:"Jono",
    gender:"M",
    domisili:"Surabaya"
},{
    _id: "u2",
    name:"Joni",
    gender:"M",
    domisili:"Malang"
},{
    _id: "u3",
    name:"Joko",
    gender:"M",
    domisili:"Semarang"
},{
    _id: "u4",
    name:"Jojo",
    gender:"F",
    domisili:"Sidoarjo"
},{
    _id: "u5",
    name:"Joji",
    gender:"M",
    domisili:"Tangerang"
}]);

db.transaction.insertMany([{
_id: new ObjectId(),
user_id:"u1",
dibeli: [{
    barang_id:"goreng1",
    qty:1
},{
    barang_id:"kuah1",
    qty:1
},{
    barang_id:"ramen2",
    qty:2
}],
totalqty: 4
},{
    _id: new ObjectId(),
    user_id:"u2",
    dibeli: [{
        barang_id:"goreng1",
        qty:2
    },{
        barang_id:"kuah1",
        qty:1
    },{
        barang_id:"ramen2",
        qty:1
    }],
    totalqty: 4
    },{
        _id: new ObjectId(),
        user_id:"u3",
        dibeli: [{
            barang_id:"goreng1",
            qty:1
        },{
            barang_id:"kuah1",
            qty:2
        },{
            barang_id:"ramen2",
            qty:1
        }],
        totalqty: 4
}])

db.products.find({
    "products": {
      "$elemMatch": {
        "_id": ObjectId("62384bec7d92848ceeb62a42"),
        "default": true
      }
    }
  }, {
    "products.$._id": 2 // "accounts.$": 1 also works
}).pretty()

  db.user.find({user: {$exists:true}}).forEach( function(x) {
    x.user = ObjectId(x.user);
    db.booking.update({_id: x._id}, {$set: {user: x.user}});
});

db.products.find( { _id: ObjectId("62384bec7d92848ceeb62a42") } )

db.products.find( {"barang._id": "62384bec7d92848ceeb62a42" } ).pretty()

db.products.find(
    { barang: { $elemMatch: { barang_id: "62384bec7d92848ceeb62a42" } } }
 )

 db.products.aggregate([
    { $unwind: "$barang"},
    {
        $match: {
            "barang._id": "p3"
        }
    }
])

db.books.insertMany([{
    _id:1,
    title:"Wawasan Pancasila",
    authorID:1,
    publisherID:1,
    price:71200,
    stats:{
        page:324,
        weight:300
    },
    publishedAt: new Date ("2018-10-01"),
    catagory:[
        "social",
        "politic"
    ]},{
    _id:2,
    title:"Mata Air Keteladanan",
    authorID:1, 
    publisherID:2,
    price:106250,
    stats:{
        page:672,
        weight:650
    },
    publishedAt: new Date ("2017-09-01"),
    catagory:[
        "social",
        "politic"
    ]},{
    _id:3,
    title:"Revolusi Pancasila",
    authorID:1,
    publisherID:1,
    price:54400,
    stats:{
        page:220,
        weight:500
    },
    publishedAt: new Date ("2015-05-01"),
    catagory:[
        "social",
        "politic"
    ]},{
    _id:4,
    title:"Self Driving",
    authorID:2,
    publisherID:1,
    price:58650,
    stats:{
        page:286,
        weight:300
    },
    publishedAt: new Date ("2018-105-01"),
    catagory:[
        "self-development"
    ]},{
    _id:5,
    title:"Self Distruption",
    authorID:2,
    publisherID:2,
    price:83300,
    stats:{
        page:400,
        weight:800
    },
    publishedAt: new Date ("2018-05-01"),
    catagory:[
        "self-development"
    ]}
])

db.authors.insertMany([{
    _id: 1,
    firstName:"Yudi",
    lastName:"Latif",
},{
    _id:2,
    firstName:"Rhenald",
    lastName:"Kasali"
}])

db.publishers.insertMany([{
    _id: 1,
    publisherName:"Expose",
},{
    _id:2,
    publisherName:"Mizan"
}])


db.books.findOne({"authorID":'1'})
db.books.find({authorID:1},{title:1,price:1})

db.books.aggregate([
    { $lookup:
        {
           from: "authors",
           localField: "authorID",
           foreignField: "_id",
           as: "authors"
        }
    }
]).pretty();

db.books.aggregate([ { 
    $group: { 
        _id: "$authorID",
        totalpages: { 
            $sum: "$stats.page" 
        } 
    } 
} ] )

db.authors.aggregate([

    // Join with user_info table
    {
        $lookup:{
            from: "books",       // other table name
            localField: "_id",   // name of users table field
            foreignField: "_id", // name of userinfo table field
            as: "books"         // alias for userinfo table
        }
    },
    {   $unwind:"$books" },     // $unwind used for getting data in object or for one record only

    // Join with user_role table
    {
        $lookup:{
            from: "publishers", 
            localField: "publisherID", 
            foreignField: "_id",
            as: "publisher"
        }
    },
    {   $unwind:"$publisher" },

    // define some conditions here 
    {
        $match:{
            $and:[{"userName" : "admin"}]
        }
    },

    // define which fields are you want to fetch
    {   
        $project:{
            _id : 1,
            email : 1,
            userName : 1,
            userPhone : "$user_info.phone",
            role : "$user_role.role",
        } 
    }
]);

db.books.aggregate([
    {
        $lookup:{
            from: "authors",      
            localField: "authorID",  
            foreignField: "_id", 
            as: "authors"       
        }
    },
    {   $unwind:"$authors" },

    {
        $lookup:{
            from: "publishers", 
            localField: "publisherID", 
            foreignField: "_id",
            as: "publisher"
        }
    },
    {   $unwind:"$publisher" }

])

db.authors.aggregate([
    { $lookup:
        {
           from: "books",
           localField: "_id",
           foreignField: "_id",
           as: "book"
        }
    }
]).pretty();

db.books.aggregate( [ { $project : { title : 1 , price : 1 } } ] )

db.authors.aggregate([
    {
      "$lookup": {
        "from": "books",
        "localField": "_id",
        "foreignField": "_id",
        "as": "listbooks"
      }
    },
    {
      "$unwind": "$listbooks"
    },
    {
      "$lookup": {
        "from": "publishers",
        "localField": "_id",
        "foreignField": "_id",
        "as": "publishers"
      }
    },
   {
      $replaceRoot: { newRoot: { $mergeObjects: [ { $arrayElemAt: [ "$listbook", 0 ] }, "$$ROOT" ] } }
   },
    {
      $project: {
        _id: 0,
        _id: {
            $concat: ['$firstName', ' ', '$lastName']
        },
        price:1
      }
    }
])

