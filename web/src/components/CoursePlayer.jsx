import { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import CourseStep from './CourseStep';
import Navigation from './Navigation';
import ProgressBar from './ProgressBar';
import { courseSteps, courses } from '../data/courseData';

function CourseComplete({ course, onBackToList }) {
    return (
        <div className="text-center py-10 px-5">
            <h2 className="text-2xl font-bold text-tg-button mb-4">🎉 Поздравляем!</h2>
            <p className="mb-3 text-tg-text">Вы успешно завершили курс "{course.title}"!</p>
            <p className="mb-6 text-tg-text">Теперь вы знаете основы программирования на JS.</p>

            <div className="flex flex-col sm:flex-row gap-3 justify-center">
                <button
                    className="bg-tg-button text-tg-button-text px-6 py-3 rounded-lg font-medium hover:opacity-90 transition-opacity"
                    onClick={() => window.location.reload()}
                >
                    Пройти курс заново
                </button>
                <button
                    className="bg-tg-secondary-bg text-tg-text border border-tg-section-separator px-6 py-3 rounded-lg font-medium hover:bg-tg-bg transition-colors"
                    onClick={onBackToList}
                >
                    ← Выбрать другой курс
                </button>
            </div>
        </div>
    );
}

export default function CoursePlayer({ user }) {
    const { courseId, stepNumber } = useParams();
    const navigate = useNavigate();

    const [currentStep, setCurrentStep] = useState(
        stepNumber ? parseInt(stepNumber) : 1
    );
    const [answers, setAnswers] = useState({});
    const [showResults, setShowResults] = useState({});

    const course = courses.find(c => c.id === courseId);

    if (!course) {
        return (
            <div className="max-w-2xl mx-auto p-5 min-h-screen">
                <div className="text-center py-10 px-5">
                    <h2 className="text-2xl font-bold text-tg-destructive mb-4">❌ Курс не найден</h2>
                    <p className="mb-6 text-tg-text">Запрашиваемый курс не существует или ещё не готов.</p>
                    <button
                        className="bg-tg-button text-tg-button-text px-6 py-3 rounded-lg font-medium hover:opacity-90 transition-opacity"
                        onClick={() => navigate('/')}
                    >
                        ← Вернуться к списку курсов
                    </button>
                </div>
            </div>
        );
    }

    const handleAnswer = (answerIndex) => {
        setAnswers({
            ...answers,
            [currentStep]: answerIndex
        });

        // Показать результат через секунду
        setTimeout(() => {
            setShowResults({
                ...showResults,
                [currentStep]: true
            });
        }, 500);
    };

    const goToNext = () => {
        if (currentStep < courseSteps.length) {
            const nextStep = currentStep + 1;
            setCurrentStep(nextStep);
            navigate(`/course/${courseId}/step/${nextStep}`, { replace: true });
        }
    };

    const goToPrev = () => {
        if (currentStep > 1) {
            const prevStep = currentStep - 1;
            setCurrentStep(prevStep);
            navigate(`/course/${courseId}/step/${prevStep}`, { replace: true });
        }
    };

    const handleBackToList = () => {
        navigate('/');
    };

    const currentStepData = courseSteps[currentStep - 1];
    const canGoNext = showResults[currentStep];

    if (currentStep > courseSteps.length) {
        return (
            <div className="max-w-2xl mx-auto p-5 min-h-screen">
                <CourseComplete course={course} onBackToList={handleBackToList} />
            </div>
        );
    }

    return (
        <div className="max-w-2xl mx-auto p-5 min-h-screen">
            <div className="flex flex-col sm:flex-row items-center gap-4 mb-5 pb-4 border-b border-tg-section-separator">
                <button
                    className="bg-tg-secondary-bg border border-tg-section-separator px-3 py-2 rounded-lg text-sm hover:bg-tg-bg transition-colors self-start"
                    onClick={handleBackToList}
                >
                    ← Курсы
                </button>
                <div className="flex items-center gap-3 flex-1 text-center sm:text-left">
          <span
              className="w-10 h-10 rounded-xl flex items-center justify-center text-xl"
              style={{ backgroundColor: course.color }}
          >
            {course.icon}
          </span>
                    <div>
                        <h2 className="text-xl font-semibold text-tg-text">{course.title}</h2>
                        {user && <p className="text-sm text-tg-hint">Удачи, {user.first_name}!</p>}
                    </div>
                </div>
            </div>

            <ProgressBar current={currentStep} total={courseSteps.length} />

            <CourseStep
                step={currentStepData}
                onAnswer={handleAnswer}
                selectedAnswer={answers[currentStep]}
                showResult={showResults[currentStep]}
            />

            <Navigation
                currentStep={currentStep}
                totalSteps={courseSteps.length}
                onPrev={goToPrev}
                onNext={goToNext}
                canGoNext={canGoNext}
            />
        </div>
    );
}