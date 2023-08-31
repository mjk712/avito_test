create table if not exists Segment (              
    id                  BIGSERIAL primary key,
    name               varchar not null
);

create table if not exists avito_user (
    id                  BIGSERIAL primary key
);

create table if not exists user_segment (
    id                  BIGSERIAL primary key,
    UserId              int not null,
    SegmentId          int not null,
    

    foreign key (UserId) references avito_user(id),
    foreign key (SegmentId) references Segment(id)
);