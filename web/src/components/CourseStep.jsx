export default function CourseStep({ step, onAnswer, selectedAnswer, showResult }) {
    const handleAnswerSelect = (answerIndex) => {
        if (!showResult) {
            onAnswer(answerIndex);
        }
    };

    return (
        <div className="bg-tg-secondary-bg rounded-xl p-6 mb-5 border border-tg-section-separator animate-slide-up">
            <div className="flex items-center mb-4">
                <div className="bg-tg-button text-tg-button-text w-8 h-8 rounded-full flex items-center justify-center font-bold text-sm">
                    {step.id}
                </div>
                <h2 className="text-xl font-semibold ml-3 flex-1 text-tg-text">
                    {step.title}
                </h2>
            </div>

            <div className="bg-tg-bg p-4 rounded-lg mb-4 border-l-4 border-tg-link">
                <h3 className="text-base font-medium mb-2">📖 Теория</h3>
                <p className="text-tg-text leading-relaxed">{step.theory}</p>
            </div>

            <div className="bg-gray-800 text-gray-300 p-4 rounded-lg my-3 font-mono text-sm overflow-x-auto">
                <pre className="whitespace-pre-wrap">{step.code}</pre>
            </div>

            <div className="bg-tg-bg p-4 rounded-lg mt-4">
                <h3 className="text-base font-medium mb-2">❓ Проверка знаний</h3>
                <div className="mb-3 font-medium text-tg-text">
                    <strong>{step.quiz.question}</strong>
                </div>

                <div className="space-y-2">
                    {step.quiz.options.map((option, index) => (
                        <div
                            key={index}
                            className={`
                bg-tg-secondary-bg border-2 border-transparent rounded-lg p-3 cursor-pointer 
                transition-all duration-200 select-none
                hover:border-tg-link
                ${selectedAnswer === index ? 'bg-tg-button text-tg-button-text border-tg-button' : ''}
                ${showResult && index === step.quiz.correct ? 'bg-green-500 text-white border-green-500' : ''}
                ${showResult && selectedAnswer === index && index !== step.quiz.correct ? 'bg-red-500 text-white border-red-500' : ''}
              `}
                            onClick={() => handleAnswerSelect(index)}
                        >
                            {option}
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}