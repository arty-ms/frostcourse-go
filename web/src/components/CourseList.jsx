import { Link } from 'react-router-dom';
import { courses } from '../data/courseData';

function CourseCard({ course }) {
    const CardContent = () => (
        <>
            <div
                className="w-12 h-12 rounded-xl flex items-center justify-center text-2xl flex-shrink-0"
                style={{ backgroundColor: course.color }}
            >
                <span>{course.icon}</span>
            </div>

            <div className="flex-1">
                <div className="flex items-center gap-2 mb-2">
                    <h3 className="text-lg font-semibold text-tg-text">{course.title}</h3>
                    {course.comingSoon && (
                        <span className="bg-tg-hint text-white text-xs px-2 py-1 rounded-full font-medium">
              Скоро
            </span>
                    )}
                </div>

                <p className="text-tg-hint text-sm mb-3 leading-tight">
                    {course.description}
                </p>

                <div className="flex gap-3 text-xs">
          <span className="text-tg-button bg-tg-bg px-2 py-1 rounded font-medium">
            {course.difficulty}
          </span>
                    <span className="text-tg-hint bg-tg-bg px-2 py-1 rounded font-medium">
            {course.duration}
          </span>
                    <span className="text-tg-hint bg-tg-bg px-2 py-1 rounded font-medium">
            {course.steps} шагов
          </span>
                </div>
            </div>

            {!course.comingSoon && (
                <div className="text-xl text-tg-button flex-shrink-0">
                    →
                </div>
            )}
        </>
    );

    if (course.comingSoon) {
        return (
            <div className={`
        bg-tg-secondary-bg border border-tg-section-separator rounded-xl p-5 
        flex items-center gap-4 opacity-60 cursor-not-allowed animate-slide-up
      `}>
                <CardContent />
            </div>
        );
    }

    return (
        <Link
            to={`/course/${course.id}`}
            className={`
        bg-tg-secondary-bg border border-tg-section-separator rounded-xl p-5 
        flex items-center gap-4 cursor-pointer transition-all duration-200 
        hover:shadow-lg hover:-translate-y-0.5 animate-slide-up block
      `}
        >
            <CardContent />
        </Link>
    );
}

export default function CourseList({ user }) {
    return (
        <div className="max-w-2xl mx-auto p-5 min-h-screen">
            <div className="text-center mb-5 animate-fade-in">
                <h1 className="text-2xl font-bold mb-2">🧊 Frostcourse</h1>
                {user ? (
                    <p className="text-tg-hint">Привет, {user.first_name}! Выберите курс для изучения</p>
                ) : (
                    <p className="text-tg-hint">Выберите курс для изучения</p>
                )}
            </div>

            <div className="grid gap-4 mb-6 md:grid-cols-1">
                {courses.map((course, index) => (
                    <div
                        key={course.id}
                        style={{ animationDelay: `${index * 0.1}s` }}
                    >
                        <CourseCard course={course} />
                    </div>
                ))}
            </div>

            <div className="text-center p-5 text-tg-hint text-sm">
                💡 Больше курсов появится совсем скоро!
            </div>
        </div>
    );
}