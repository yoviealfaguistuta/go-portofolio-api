-- Create tables
CREATE TABLE certificates (
    id              BIGSERIAL       PRIMARY KEY,
    images          TEXT            NOT NULL,
    title           VARCHAR (255)   NOT NULL,
    organization      VARCHAR (255)   NOT NULL,
    publish         VARCHAR (100)   NOT NULL,
    credentials     VARCHAR (250)   NOT NULL DEFAULT '-',
    urls            VARCHAR (250)   NOT NULL DEFAULT '-',
    created_at      TIMESTAMP       NOT NULL,
    updated_at      TIMESTAMP       NOT NULL
);

INSERT INTO certificates (id, images, title, organization, publish, credentials, urls, created_at, updated_at) VALUES 
(1, '1645690383_Dicoding - Android Developer Roadmap.png', 'Dicoding - Android Developer Roadmap', 'Dicoding', '19 Juli 2020', '-', 'https://dicoding.com/users/yoviealfa80/events', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(2, '1645690412_KOMINFO - Intro to Fullstack Developer.png', 'KOMINFO - Intro to Fullstack Developer', 'KOMINFO', '31 Mei 2021', '-', '-', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(3, '1645690562_Dicoding - AWS Cloud Backend.png', 'Dicoding - Cloud Practitioner Essentials', 'Dicoding', '19 Agustus 2021', 'QLZ9144K7P5D', 'https://dicoding.com/certificates/QLZ9144K7P5D', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(4, '1645690482_Dicoding - Belajar Prinsip Pemrograman SOLID.png', 'Dicoding - Belajar Prinsip Pemrograman SOLID', 'Dicoding', '19 Agustus 2021', '1OP8LNN02ZQK', 'https://dicoding.com/certificates/1OP8LNN02ZQK', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(5, '1645690503_Progate - Python Course.png', 'Progate - Python Course', 'Progate', '20 Juni 2021', '3a5570f3qv06zz', 'https://progate.com/course_certificate/3a5570f3qv06zz', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(6, '1645690503_Progate - Python Course.png', 'Progate - Python Course', 'Progate', '20 Juni 2021', '3a5570f3qv06zz', 'https://progate.com/course_certificate/3a5570f3qv06zz', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(7, '1645690555_Progate - Command Line Course.png', 'Progate - Command Line Course', 'Progate', '28 Agustus 2021', '94da83eaqyjyyk', 'https://progate.com/course_certificate/94da83eaqyjyyk', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(8, '1645690547_Progate - GIT Course.png', 'Progate - GIT Course', 'Progate', '28 Agustus', 'c0b36072qyjyd5', 'https://progate.com/course_certificate/c0b36072qyjyd5', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(9, '1645690531_Progate - HTML & CSS Course.png', 'Progate - HTML & CSS Course', 'Progate', '20 Juni 2021', '5f0a274bqv07l8', 'https://progate.com/course_certificate/5f0a274bqv07l8', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(10, '1645690522_Progate - Web Development Path (Node.js).png', 'Progate - Web Development Path (Node.js)', 'Progate', '20 Juni 2021', '2ca68de3qv05pg', 'https://progate.com/path_certificate/2ca68de3qv05pg', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(11, '1645690512_Progate - Javascript Course.png', 'Progate - Javascript Course', 'Progate', '3 Oktober 2020', '7abb108aqhlv1q', 'https://progate.com/course_certificate/7abb108aqhlv1q', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(12, '1645690976_Teknokrat - Best UX Design.png', 'Teknokrat - Best UX Design', 'Universitas Teknokrat', '23 Januari 2020', '-', '-', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(13, '1645690539_Teknokrat - Best UX Design Team.png', 'Teknokrat - Best UX Design Team', 'Universitas Teknokrat', '23 Januari 2020', '-', '-', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(14, '1645690539_Teknokrat - Best UX Design Team.png', 'Teknokrat - Best UX Design Team', 'Universitas Teknokrat', '23 Januari 2020', '-', '-', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(15, 'react-redux-course.png', 'React + Redux Course', 'Sololearn', '11 July 2022', '1097-18790924', 'https://www.sololearn.com/Certificate/1097-18790924/jpg', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(16, 'React - The Complete Guide with React Hook Redux.png', 'React - The Complete Guide with React Hook Redux', 'Udemy', 'Februari 2022', 'UC-8270e54e-5d5a-47ad-bc58-7cac02cd8f60', 'https://www.udemy.com/certificate/UC-8270e54e-5d5a-47ad-bc58-7cac02cd8f60', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(17, 'Java Course.png', 'Java Course', 'Sololearn', 'Juni 2020', '1068-18790924', 'https://www.sololearn.com/Certificate/1068-18790924/jpg/', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(18, 'SQL Course.png', 'SQL Course', 'Sololearn', 'Juni 2020', '1068-18790924', 'https://www.sololearn.com/Certificate/1068-18790924/jpg/', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(19, 'React - The Complete Guide with React Hook Redux.png', 'React - The Complete Guide with React Hook Redux', 'Udemy', 'Februari 2022', 'UC-8270e54e-5d5a-47ad-bc58-7cac02cd8f60', 'https://www.udemy.com/certificate/UC-8270e54e-5d5a-47ad-bc58-7cac02cd8f60', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(20, 'react-redux-course.jpg', 'React + Redux Course', 'Sololearn', 'Juli 2022', '1097-18790924', 'https://www.sololearn.com/Certificate/1097-18790924/jpg', '2021-08-03 18:00:19', '2021-08-03 18:00:19'),
(21, 'Sertifikat Devcode Challenge.png', 'Sertifikat Devcode Challenge', 'Skyshi', 'Agustus 2022', 'ICL480', 'https://devcode.gethired.id/job/ICL480', '2021-08-03 18:00:19', '2021-08-03 18:00:19');