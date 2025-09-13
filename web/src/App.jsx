import { useEffect, useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import CourseList from './components/CourseList';
import CoursePlayer from './components/CoursePlayer';
import './styles/main.css';

export default function App() {
    const [user, setUser] = useState(null);

    useEffect(() => {
        // Инициализация Telegram Web App
        if (window.Telegram?.WebApp) {
            window.Telegram.WebApp.ready();
            window.Telegram.WebApp.expand();

            const tgUser = window.Telegram.WebApp.initDataUnsafe?.user;
            if (tgUser) {
                setUser(tgUser);
            }
        }
    }, []);

    return (
        <Router>
            <Routes>
                <Route
                    path="/"
                    element={<CourseList user={user} />}
                />
                <Route
                    path="/course/:courseId"
                    element={<CoursePlayer user={user} />}
                />
                <Route
                    path="/course/:courseId/step/:stepNumber"
                    element={<CoursePlayer user={user} />}
                />
            </Routes>
        </Router>
    );
}