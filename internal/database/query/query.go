package query

var GetUserSegments = `
SELECT u.id AS userId, s.name AS segmentName
FROM avito_user u 
LEFT JOIN user_segment us ON us.userId = u.id
LEFT JOIN segment s ON s.id = us.segmentid
WHERE u.id = $1;`

var DeleteUserFromSegment = `DELETE FROM user_segment WHERE userid= $1 AND segmentid = $2;`

var AddUserIntoSegment = `
INSERT INTO user_segment(userId,segmentId)
VALUES($1,$2);`

var AddSegment = `
INSERT INTO segment(name)
VALUES($1);`

var DeleteSegment = `DELETE FROM segment WHERE name= $1;`

var GetSegmentIdByName = `SELECT id FROM segment WHERE name=$1;`

var DeleteSegmentFromUserSegment = `DELETE FROM user_segment WHERE segmentid = $1;`
