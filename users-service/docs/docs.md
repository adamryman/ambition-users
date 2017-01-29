
<style type="text/css">

body{
    font-family      : helvetica, arial, freesans, clean, sans-serif;
    color            : #003269;
    background-color : #fff;
    border-color     : #999999;
    border-width     : 2px;
    line-height      : 1.5;
    margin           : 2em 3em;
    text-align       :left;
    font-size        : 16px;
    padding          : 0 100px 0 100px;

    width         : 1024px;
    margin-top    : 0px;
    margin-bottom : 2em;
    margin-left   : auto;
    margin-right  : auto;
}

h1 {
    font-family : 'Gill Sans Bold', 'Optima Bold', Arial, sans-serif;
    color       : #577AD3;
    font-weight : 400;
    font-size   : 48px;
}
h2{
    margin-bottom : 1em;
    padding-top   : 0.5em;
    color         : #003269;
    font-size     : 36px;
}
h3{
    border-bottom : 1px dotted #aaa;
    color         : #4660A4;
    font-size     : 30px;
}
h4 {
    font-size: 24px;
}
h5 {
    font-size: 18px;
}
code {
    font-family      : Consolas, "Inconsolata", Menlo, Monaco, Lucida Console, Liberation Mono, DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace, serif; /* Taken from the stackOverflow CSS*/
    background-color : #f5f5f5;
    border           : 1px solid #e1e1e8;
}


pre {
    display          : block;
    background-color : #f5f5f5;
    border           : 1px solid #ccc;
    padding          : 3px 3px 3px 3px;
}
pre code {
    white-space      : pre-wrap;
    padding          : 0;
    border           : 0;
    background-color : code;
}

table {
	border-collapse: collapse; border-spacing: 0;
	width: 100%;
	margin-bottom : 3em;
}
td, th {
	vertical-align: top;
	padding: 4px 10px;
	border: 1px solid #9BC3EB;
}
tr:nth-child(even) td, tr:nth-child(even) th {
	background: #EBF4FE;
}
th:nth-child(4) {
	width: auto;
}

</style>
# users

## users.proto

### Messages

<a name="User"></a>

#### User

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| ID | TYPE_INT64 | 1 | ID = 0 for create Id is immutable |
| Info | [UserInfo](#UserInfo) | 2 |  |
| Trello | [TrelloInfo](#TrelloInfo) | 9 |  |

<a name="UserInfo"></a>

#### UserInfo

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| Username | TYPE_STRING | 1 |  |
| Email | TYPE_STRING | 2 |  |
| Hash | TYPE_STRING | 3 |  |
| Salt | TYPE_STRING | 4 |  |

<a name="TrelloInfo"></a>

#### TrelloInfo

TrelloInfo comes from trello, this information is used by
 ambition-rello

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| ID | TYPE_STRING | 1 |  |
| AvatarHas | TYPE_STRING | 2 |  |
| FullName | TYPE_STRING | 3 |  |
| Initials | TYPE_STRING | 4 |  |
| Username | TYPE_STRING | 5 |  |
| WebhookURL | TYPE_STRING | 6 |  |

### Services

#### Users

Users stores user information. It assumes all requests are authenticated.

| Method Name | Request Type | Response Type | Description|
| ---- | ---- | ------------ | -----------|
| CreateUser | User | User | Create takes information about a user and creates an entry
 It returns that User with it's ID populated |
| ReadUser | User | User | ReadUser takes some information about a user and tries to find the
 user with that information
 Accepted values: ID, Info.Username, Info.email, Trello.Id, and
 Trello.Username |
| UpdateUser | User | User | UpdateUser requires an ID, which can be obttained from ReadUser.
 All other non-zerp values will be updated |
| DeleteUser | User | User | DeleteUser requires an ID, as a Read before a delete is a good idea.
 It will return User with just thre requested ID on success |

#### Users - Http Methods

