// MongoDB initialization script
// This runs on first container start when the database is created

db = db.getSiblingDB('posts_service');

// Create collections with schema validation
db.createCollection('posts');
db.createCollection('authors');

// Create indexes for posts collection
db.posts.createIndex({ "slug": 1 }, { unique: true });
db.posts.createIndex({ "author_id": 1 });
db.posts.createIndex({ "created_at": -1 });

// Create indexes for authors collection
db.authors.createIndex({ "email": 1 }, { unique: true });

// Seed an initial author for development
db.authors.insertOne({
    name: "Dev Author",
    email: "dev@example.com"
});

print("Database initialized with collections, indexes, and seed data.");
