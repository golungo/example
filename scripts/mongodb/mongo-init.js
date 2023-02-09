/*
  MongoDB initial setup

  Needed only in Dev mode
*/
print('Start #################################################################');

const MOCK_DATA_FOLDER = '/tmp/mockData';
const DEFAULT_USER = "test";
const DEFAULT_PASSWORD = "testpassword";
const DEFAULT_DB_NAME = "test";


_projectDb = db.getSiblingDB(DEFAULT_DB_NAME);
_projectDb.createUser(
  {
    user: DEFAULT_USER,
    pwd: DEFAULT_PASSWORD,
    roles: [{ role: 'readWrite', db: DEFAULT_DB_NAME }],
  },
);  

_projectDb.auth({
  user: DEFAULT_USER,
  pwd: DEFAULT_PASSWORD
})

_dataDb = new Mongo().getDB(DEFAULT_DB_NAME);

const _RawItemsData = fs.readFileSync(`${MOCK_DATA_FOLDER}/items.json`);
const _RawCategoryData = fs.readFileSync(`${MOCK_DATA_FOLDER}/categories.json`);

const _CategoryData = JSON.parse(_RawCategoryData).map(category => {
  return {
    ...category,
    _id: ObjectId()
  }
})

const _ItemsData = JSON.parse(_RawItemsData).map(item => {
  const { kind, ...plainItem } = item;
  const category = _CategoryData.find(_c => _c.title === kind);

  return {
    ...plainItem,
    categoryId: category._id,
    _id: ObjectId()
  }
})

_dataDb.createCollection("categories", { capped: false });
_dataDb.categories.insertMany(_CategoryData)

_dataDb.createCollection("items", { capped: false });
_dataDb.items.insertMany(_ItemsData)

print('END #################################################################');